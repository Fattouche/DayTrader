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

# Uncomment to switch back to singleton
# @singleton


class Socket():
    def __init__(self):
        self.socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.socket.connect(
            (settings.QUOTE_SERVER_HOST, settings.QUOTE_SERVER_PORT))


class Stock:
    def __init__(self, symbol, price):
        self.symbol = symbol
        self.price = Decimal(price)

    def check_sell_trigger(self):
        sell_trigger = SellTrigger.objects.filter(
            sell__stock_symbol=self.symbol, active=True, committed=False)
        for trigger in sell_trigger:
            trigger.check_validity(self.price)

    def check_buy_trigger(self):
        buy_trigger = BuyTrigger.objects.filter(
            buy__stock_symbol=self.symbol, active=True, committed=False)
        for trigger in buy_trigger:
            trigger.check_validity(self.price)

    def verify_triggers(self):
        self.check_sell_trigger()
        self.check_buy_trigger()

    def execute_quote_request(self, user_id):
        request = "{},{}\r".format(self.symbol, user_id)
        socket = Socket()
        socket.socket.send(request.encode())
        data = socket.socket.recv(1024)
        response = data.decode().split(",")  # log the timestamp etc from this response
        quote_price = response[0]
        self.price = Decimal(quote_price)

        # TODO(cailan): deal with server name
        # TODO(cailan): deal with transaction number
        AuditLogger.log_quote_server_event('BEAVER_1', 1, quote_price,
                                           self.symbol, user_id, response[3], response[4])

    @classmethod
    def quote(cls, symbol, user_id):
        stock = cache.get(symbol)
        if(stock is None):
            stock = cls(symbol=symbol, price=0)
            stock.execute_quote_request(user_id=user_id)
            cache.set(symbol, stock, 60)
            django_rq.enqueue(stock.verify_triggers)
        return stock


class User(models.Model):
    user_id = models.CharField(max_length=100, primary_key=True)
    name = models.TextField(max_length=100)
    balance = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    password = models.CharField(max_length=50)

    @classmethod
    def get(cls, user_id):
        ##user = cache.get(user_id)
        # if(user is None):
        user = cls.objects.get(user_id=user_id)
        user.sell_stack = []
        user.buy_stack = []
        cache.set(user.user_id, user)
        return user

    def perform_buy(self, symbol, amount):
        buy, err = Buy.create(stock_symbol=symbol,
                              cash_amount=amount, user=self)
        if(err):
            return err
        self.buy_stack.append(buy)
        cache.set(self.user_id, self)

    def cancel_buy(self):
        buy = self.pop_from_buy_stack()
        if buy is not None:
            buy.cancel(self)

    def perform_sell(self, symbol, amount):
        user = User.get(self.user_id)
        user_stock, created = UserStock.objects.get_or_create(
            stock_symbol=symbol, user=user)
        sell, err = Sell.create(stock_symbol=symbol,
                                cash_amount=amount, user=self)
        if(err):
            return err
        self.sell_stack.append(sell)
        cache.set(self.user_id, self)

    def cancel_sell(self):
        sell = self.pop_from_sell_stack()
        if sell is not None:
            sell.cancel()

    def set_buy_amount(self, symbol, amount):
        try:
            buy_trigger = BuyTrigger.objects.get(
                user=self,
                buy__symbol=symbol,
                commited=False
            )
            buy_trigger.update_cash_amount(amount)
        except:
            buy, err = Buy.create(stock_symbol=symbol,
                                  cash_amount=amount, user=self)
            if(err):
                return err
            buy.save()
            buy_trigger = BuyTrigger.objects.create(
                user=self,
                buy=buy
            )

    def set_sell_amount(self, symbol, amount):
        try:
            sell_trigger = SellTrigger.objects.get(
                user=self,
                sell__symbol=symbol,
                commited=False
            )
            sell_trigger.update_cash_amount(amount)
        except:
            sell, err = Sell.create(
                stock_symbol=symbol, cash_amount=amount, user=self)
            if(err):
                return err
            sell.save()
            sell_trigger = SellTrigger.objects.create(
                user=self,
                sell=sell
            )

    def set_buy_trigger(self, symbol, price):
        try:
            buy_trigger = BuyTrigger.objects.get(
                buy__stock_symbol=symbol, user=self.user_id, committed=False)
            buy_trigger.update_trigger_price(price)
        except ObjectDoesNotExist:
            return "Trigger requires a buy amount first, please make one"

    def set_sell_trigger(self, symbol, price):
        try:
            sell_trigger = SellTrigger.objects.get(
                sell__stock_symbol=symbol, user=self.user_id, committed=False)
            sell_trigger.update_trigger_price(price)
        except ObjectDoesNotExist:
            return "Trigger requires a sell amount first, please make one"

    def cancel_set_buy(self, symbol):
        try:
            buy_trigger = BuyTrigger.objects.get(
                buy__stock_symbol=symbol, user=self.user_id, committed=False)
            buy_trigger.cancel()
        except ObjectDoesNotExist:
            return "buy trigger not found"

    def cancel_set_sell(self, symbol):
        try:
            sell_trigger = SellTrigger.objects.get(
                sell__stock_symbol=symbol, user=self.user_id, committed=False)
            sell_trigger.cancel()
        except ObjectDoesNotExist:
            return "sell trigger not found"

    def update_balance(self, change):
        self.balance += change
        cache.set(self.user_id, self)
        self.save()
        # TODO(cailan): fix server_name and transaction_num
        action = 'add' if change >= 0 else 'remove'
        AuditLogger.log_account_transaction('BEAVER_1', 1, action, self.user_id,
                                            abs(change))

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
    def create(cls, stock_symbol, cash_amount, user):
        stock = Stock.quote(stock_symbol, user.user_id)
        sell = cls(user=user, stock_symbol=stock_symbol)
        err = sell.update_price(stock.price)
        if(err):
            return None, err
        err = sell.update(cash_amount)
        if(err):
            return None, err
        return sell, None

    def update_price(self, stock_price):
        self.cancel()
        user_stock, created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        self.stock_sold_amount = min(Decimal(
            self.intended_cash_amount)//Decimal(stock_price),
            user_stock.amount)
        if(self.stock_sold_amount <= 0):
            return "Update trigger price failed"
        self.actual_cash_amount = Decimal(
            self.stock_sold_amount)*Decimal(stock_price)
        self.timestamp = time.time()
        self.sell_price = stock_price
        user_stock.update_amount(self.stock_sold_amount*-1)

    def update_cash_amount(self, amount):
        stock_price = Stock.quote(self.stock_symbol, self.user.user_id)
        user_stock, created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        stock_sold_amount = Decimal(amount)//Decimal(stock_price)
        if(stock_sold_amount > user_stock.amount):
            return "Not enough stock, have {0} need {1}".format(stock_sold_amount, user_stock.amount)
        self.intended_cash_amount = amount

    def commit(self, user):
        user.update_balance(self.actual_cash_amount)
        self.save()

    def cancel(self):
        user_stock, created = UserStock.objects.get_or_create(
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
    def create(cls, stock_symbol, cash_amount, user):
        stock = Stock.quote(stock_symbol, user.user_id)
        buy = cls(user=user, stock_symbol=stock_symbol,
                  purchase_price=stock.price)
        err = buy.update_cash_amount(cash_amount)
        if(err):
            return None, err
        buy.update_price(stock.price)
        buy.timestamp = time.time()
        return buy, None

    def update_cash_amount(self, amount):
        if amount > self.user.balance:
            return None, "Not enough balance, have {0} need {1}".format(self.user.balance, amount)
        updated_amount = (self.intended_cash_amount - amount)
        self.user.update_balance(updated_amount)
        self.intended_cash_amount = abs(updated_amount)

    def update_price(self, stock_price):
        self.stock_price = stock_price
        self.stock_bought_amount = self.intended_cash_amount//self.stock_price
        self.actual_cash_amount = self.stock_bought_amount*self.stock_price

    def cancel(self, user):
        user.update_balance(self.intended_cash_amount)

    def commit(self):
        user_stock, created = UserStock.objects.get_or_create(
            user=self.user, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_bought_amount)
        self.save()


class SellTrigger(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    sell = models.ForeignKey(Sell, on_delete=models.CASCADE)
    # Trigger should not be used anymore
    committed = models.BooleanField(default=False)
    # Cash amount and trigger price are set.
    active = models.BooleanField(default=False)

    def check_validity(self, price):
        if(self.sell.sell_price <= price):
            self.sell.commit(self.user)
            self.committed = True
            self.save()

    def update_cash_amount(self, amount):
        self.sell.update_cash_amount(amount)
        self.save()

    def update_trigger_price(self, price):
        err = self.sell.update_price(price)
        if(err is None):
            self.active = True
        self.save()
        return err

    def cancel(self):
        self.sell.cancel()
        self.committed = True
        self.active = False
        self.save()


class BuyTrigger(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    buy = models.ForeignKey(Buy, on_delete=models.CASCADE)
    # Trigger should not be used anymore
    committed = models.BooleanField(default=False)
    # Cash amount and trigger price are set.
    active = models.BooleanField(default=False)

    def check_validity(self, price):
        if(self.price >= price):
            self.buy.update_price(price)
            user = User.get(self.user.user_id)
            self.buy.commit()
            self.committed = True
            self.save()

    def update_cash_amount(self, amount):
        self.buy.update_cash_amount(amount)
        self.save()

    def update_trigger_price(self, price):
        self.buy.update_price(price)
        self.save()

    def cancel(self):
        self.buy.cancel(self.user)
        self.committed = True
        self.save()


def is_expired(previous_time):
    elapsed_time = time.time() - previous_time
    if(elapsed_time > 60):
        return True
    return False
