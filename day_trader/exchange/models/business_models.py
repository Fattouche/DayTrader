from django.db import models
from django.core.cache import cache
from django.core.exceptions import ObjectDoesNotExist
import time
import django_rq
from decimal import Decimal
import socket
from django.conf import settings
from exchange.audit_logging import AuditLogger


def singleton(cls, *args, **kw):
    instances = {}
    def _singleton():
        if cls not in instances:
            instances[cls] = cls(*args, **kw)
        return instances[cls]
    return _singleton

#Uncomment to switch back to singleton
#@singleton
class Socket():
    def __init__(self):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect(
                (settings.QUOTE_SERVER_HOST, settings.QUOTE_SERVER_PORT))


class Stock:
    def __init__(self, symbol, price, transaction_num):
        self.symbol = symbol
        self.price = Decimal(price)
        self.transaction_num = transaction_num

    def check_sell_trigger(self):
        sell_trigger = SellTrigger.objects.filter(
            stock_symbol=self.symbol, committed=False, active=True)
        for trigger in sell_trigger:
            trigger.check_validity(self.price, self.transaction_num)

    def check_buy_trigger(self):
        buy_trigger = BuyTrigger.objects.filter(
            stock_symbol=self.symbol, committed=False, active=True)
        for trigger in buy_trigger:
            trigger.check_validity(self.price, self.symbol, 
                self.transaction_num)

    def verify_triggers(self):
        self.check_sell_trigger()
        self.check_buy_trigger()

    def execute_quote_request(self, user_id, transaction_num):
        request = "{},{}\r".format(self.symbol, user_id)
        socket = Socket()
        socket.socket.send(request.encode())
        data=socket.socket.recv(1024)
        response=data.decode().split(",")  # log the timestamp etc from this response
        quote_price=response[0]
        self.price=Decimal(quote_price)

        # TODO(cailan): deal with server name
        AuditLogger.log_quote_server_event('BEAVER_1', transaction_num, 
            quote_price, self.symbol, user_id, response[3], response[4])

    @classmethod
    def quote(cls, symbol, user_id, transaction_num):
        stock=cache.get(symbol)
        if(stock is None):
            stock=cls(symbol, 0, transaction_num)
            stock.execute_quote_request(user_id, transaction_num)
            cache.set(symbol, stock, 60)
            django_rq.enqueue(stock.verify_triggers)
        return stock


class User(models.Model):
    user_id=models.CharField(max_length=100, primary_key=True)
    name=models.TextField(max_length=100)
    balance=models.DecimalField(max_digits=65, decimal_places=2, default=0)
    password=models.CharField(max_length=50)

    @classmethod
    def get(cls, user_id):
        user=cache.get(user_id)
        if(user is None):
            user=cls.objects.get(user_id=user_id)
            user.sell_stack=[]
            user.buy_stack=[]
            cache.set(user.user_id, user)
        return user

    def perform_buy(self, symbol, amount, transaction_num):
        if Decimal(amount) > self.balance:
            return "Not enough balance, have {0} need {1}".format(self.balance, amount)
        stock=Stock.quote(symbol, self.user_id, transaction_num)
        buy=Buy.create(stock_symbol=symbol, cash_amount=amount, 
                        stock_price=stock.price, user=self,
                        transaction_num=transaction_num)
        self.buy_stack.append(buy)
        cache.set(self.user_id, self)

    def cancel_buy(self, transaction_num):
        buy=self.pop_from_buy_stack()
        if buy is not None:
            buy.cancel(self, transaction_num)

    def perform_sell(self, symbol, amount, transaction_num):
        user=User.get(self.user_id)
        user_stock, created=UserStock.objects.get_or_create(
            stock_symbol=symbol, user=user)
        if user_stock.amount is None or user_stock.amount == 0:
            return "Not enough {0} to sell".format(symbol)
        stock = Stock.quote(symbol, self.user_id, transaction_num)
        sell = Sell.create(stock_symbol=symbol,
                           cash_amount=amount, stock_price=stock.price, user=self)
        self.sell_stack.append(sell)
        cache.set(self.user_id, self)

    def cancel_sell(self):
        sell = self.pop_from_sell_stack()
        if sell is not None:
            sell.cancel()

    def set_buy_amount(self, symbol, amount, transaction_num):
        buy_trigger, created = BuyTrigger.objects.get_or_create(
            stock_symbol=symbol,
            user=self,
            defaults={'cash_amount': amount},
        )
        self.update_balance(amount*-1, transaction_num)
        if not created:
            buy_trigger.update_cash_amount(amount)

    def set_sell_amount(self, symbol, amount):
        sell_amount_set = False
        try:
            user_stock = UserStock.objects.get(user=self.user_id, stock_symbol=symbol)
            stock = Stock.quote(symbol, self.user_id)
            stock_amount = user_stock.amount // stock.price
            if(stock_amount > 0):
                sell_trigger, created = SellTrigger.objects.get_or_create(
                    stock_symbol=symbol,
                    user=self,
                    defaults={'cash_amount': amount},
                )

                if not created:
                    sell_trigger.update_cash_amount(amount)
                
                sell_amount_set = True
        except:
            pass

        return sell_amount_set

    def set_buy_trigger(self, symbol, price):
        try:
            buy_trigger = BuyTrigger.objects.get(
                stock_symbol=symbol, user=self.user_id)
            if(buy_trigger.cash_amount > 0 and buy_trigger.cash_amount >= price):
                buy_trigger.update_trigger_price(price)
            else:
                buy_trigger = None
        except ObjectDoesNotExist:
            buy_trigger = None

        return buy_trigger is not None

    def set_sell_trigger(self, symbol, price):
        trigger_set = False
        try:
            sell_trigger = SellTrigger.objects.get(
                stock_symbol=symbol, user=self.user_id)
            if(sell_trigger.cash_amount > 0):
                trigger_set = sell_trigger.update_trigger_price(price)
        except ObjectDoesNotExist:
            pass   
        return trigger_set

    def cancel_set_buy(self, symbol, transaction_num):
        try:
            buy_trigger = BuyTrigger.objects.get(
                stock_symbol=symbol, user_id=self.user_id)
            self.update_balance(buy_trigger.cash_amount, transaction_num)
            buy_trigger.cancel()
        except ObjectDoesNotExist:
            buy_trigger = None

        return buy_trigger is not None

    def cancel_set_sell(self, symbol):
        try:
            sell_trigger = SellTrigger.objects.get(
                stock_symbol=symbol, user=self.user_id)
            sell_trigger.cancel()
        except ObjectDoesNotExist:
            sell_trigger = None

        return sell_trigger is not None

    def update_balance(self, change, transaction_num):
        self.balance += change
        cache.set(self.user_id, self)
        self.save()
        # TODO(cailan): fix server_name and transaction_num
        action = 'add' if change >= 0 else 'remove'
        AuditLogger.log_account_transaction('BEAVER_1', transaction_num, 
            action, self.user_id, abs(change))

    def pop_from_buy_stack(self):
        if self.buy_stack:
            buy = self.buy_stack.pop()
            cache.set(self.user_id, self)
            return buy
        return None

    def pop_from_sell_stack(self):
        if self.sell_stack:
            sell = self.sell_stack.pop()
            cache.set(self.user_id, self)
            return sell
        return None


class UserStock(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3)
    amount = models.PositiveIntegerField(default=0)

    def update_amount(self, change):
        self.amount += change
        self.save()


class SellTrigger(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3)
    price = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_reserved_amount = models.PositiveIntegerField(default=0)
    committed = models.BooleanField(default=False)
    active = models.BooleanField(default=False)

    def check_validity(self, price):
        if(self.price <= price):
            user = User.get(self.user.user_id)
            # TODO(isaacsahle): This should be refactored.
            actual_cash_amount = self.stock_reserved_amount * (price)
            sell = Sell.create(user=user, stock_symbol=self.stock_symbol,
                               cash_amount=actual_cash_amount, 
                               stock_price=price ,decrement_stock=False)
            sell.commit(user, transaction_num)
            self.committed = True
            self.save()

    def update_cash_amount(self, amount):
        self.cash_amount += amount
        self.save()

    def update_trigger_price(self, price):
        user_stock,created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        stock_reserve_amount = min(self.cash_amount // price, user_stock.amount)
        trigger_updated = False
        if stock_reserve_amount > 0:
            user_stock.update_amount(
                self.stock_reserved_amount - stock_reserve_amount)
            self.stock_reserved_amount = stock_reserve_amount
            self.price = price
            self.active = True
            self.save()
            trigger_updated = True
        return trigger_updated

    def cancel(self):
        user_stock,created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_reserved_amount)
        self.price = 0
        self.cash_amount = 0
        self.stock_reserved_amount = 0
        self.active = False
        self.save()


class BuyTrigger(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3)
    price = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    committed = models.BooleanField(default=False)
    active = models.BooleanField(default=False)

    def check_validity(self, price, symbol, transaction_num):
        if(self.price >= price):
            user = User.get(self.user.user_id)
            buy = Buy.create(user=user, stock_symbol=self.stock_symbol,
                             cash_amount=cash_amount, stock_price=price,
                             transaction_num=transaction_num)
            buy.commit(user)
            self.committed = True
            self.save()

    def update_cash_amount(self, amount):
        self.cash_amount += amount
        self.save()

    def update_trigger_price(self, price):
        self.price = price
        self.active = True
        self.save()

    def cancel(self):
        self.active = False
        self.price = 0
        self.cash_amount = 0
        self.save()


class Sell(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3)
    intended_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    actual_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_sold_amount = models.PositiveIntegerField(default=0)
    sell_price = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)

    @classmethod
    def create(cls, stock_symbol, cash_amount, stock_price, user, decrement_stock=True):
        sell = cls(user=user, stock_symbol=stock_symbol,
                   intended_cash_amount=cash_amount, sell_price=stock_price)
        sell.stock_sold_amount = Decimal(cash_amount)//Decimal(stock_price)
        sell.actual_cash_amount = Decimal(
            sell.stock_sold_amount)*Decimal(stock_price)
        sell.timestamp = time.time()
        if(decrement_stock):
            user_stock,created = UserStock.objects.get_or_create(
                user=user, stock_symbol=stock_symbol)
            user_stock.update_amount(sell.stock_sold_amount*-1)
        return sell

    def commit(self, user, transaction_num):
        user.update_balance(self.actual_cash_amount, transaction_num)
        self.save()

    def cancel(self):
        user_stock,created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_sold_amount)


class Buy(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3)
    intended_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    actual_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_bought_amount = models.PositiveIntegerField(default=0)
    purchase_price = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)

    @classmethod
    def create(cls, stock_symbol, cash_amount, stock_price, user,
                transaction_num):
        buy = cls(user=user, stock_symbol=stock_symbol,
                  purchase_price=stock_price)
        buy.stock_bought_amount = Decimal(cash_amount)//Decimal(stock_price)
        buy.actual_cash_amount = Decimal(buy.stock_bought_amount)*Decimal(stock_price)
        buy.timestamp = time.time()
        user.update_balance(buy.actual_cash_amount*-1, transaction_num)
        return buy

    def cancel(self, user, transaction_num):
        user.update_balance(self.actual_cash_amount, transaction_num)

    def commit(self):
        user_stock, created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_bought_amount)
        self.save()


def is_expired(previous_time):
    elapsed_time = time.time() - previous_time
    if(elapsed_time > 60):
        return True
    return False


