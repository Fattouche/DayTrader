from django.db import models
from django.core.cache import cache
from django.core.exceptions import ObjectDoesNotExist
import time
import django_rq
from decimal import Decimal
import socket


class Stock:
    def __init__(self, symbol, price):
        self.symbol = symbol
        self.price = price

    def check_sell_trigger(self):
        sell_trigger = SellTrigger.objects.filter(
            stock_symbol=self.symbol, committed=False, active=True)
        for trigger in sell_trigger:
            trigger.check_validity(self.price)

    def check_buy_trigger(self):
        buy_trigger = BuyTrigger.objects.filter(
            stock_symbol=self.symbol, committed=False, active=True)
        for trigger in buy_trigger:
            trigger.check_validity(self.price, self.symbol)

    def verify_triggers(self):
        self.check_sell_trigger()
        self.check_buy_trigger()

    def execute_quote_request(self, user_id):
        request = "{},{}\r".format(self.symbol, user_id)

        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(('localhost',4442))
        s.send(request)
        data = s.recv(1024)
        s.close()

        response = data.decode().split(",")  #log the timestamp etc from this response
        quote_price = response[0]
        self.price = quote_price

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
    balance = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    password = models.CharField(max_length=50)

    @classmethod
    def get(cls, user_id):
        user = cache.get(user_id)
        if(user is None):
            user = cls.objects.get(user_id=user_id)
            user.sell_stack = []
            user.buy_stack = []
            cache.set(user.user_id, user)
        return user

    def perform_buy(self, symbol, amount):
        if Decimal(amount) > self.balance:
            return "Not enough balance, have {0} need {1}".format(self.balance, amount)
        stock = Stock.quote(symbol, self.user_id)
        buy = Buy.create(stock_symbol=symbol,
                         cash_amount=amount, stock_price=stock.price, user=self)
        self.buy_stack.append(buy)
        cache.set(self.user_id, self)
    
    def cancel_buy(self):
        buy = self.pop_from_buy_stack()
        if buy is not None:
            buy.cancel(self)

    def perform_sell(self, symbol, amount):
        user = User.get(self.user_id)
        stock_amount = UserStock.objects.get(
            stock_symbol=symbol, user_id=user)
        if stock_amount is None or stock_amount == 0:
            return "Not enough {0} to sell".format(symbol)
        stock = Stock.quote(symbol, self.user_id)
        sell = Sell.create(stock_symbol=symbol,
                           cash_amount=amount, stock_price=stock.price, user=self)
        self.sell_stack.append(sell)
        cache.set(self.user_id, self)
    
    def cancel_sell(self):
        sell = self.pop_from_sell_stack()
        if sell is not None:
            sell.cancel(self)
 
    def set_buy_amount(self, symbol, amount):
        buy_trigger, created = BuyTrigger.objects.get_or_create(
            stock_symbol=symbol,
            user_id=self,
            defaults={'cash_amount': amount},
        )
        self.update_balance(amount*-1)
        if not created:
            buy_trigger.update_cash_amount(amount)
    
    def set_sell_amount(self, symbol, amount):
        sell_trigger, created = SellTrigger.objects.get_or_create(
            stock_symbol=symbol,
            user_id=self,
            defaults={'cash_amount': amount},
        )

        if not created:
            sell_trigger.update_cash_amount(amount)
    
    def set_buy_trigger(self, symbol, amount):
        try:
            buy_trigger = BuyTrigger.objects.get(stock_symbol=symbol,user_id=self.user_id)
            if(buy_trigger.cash_amount > 0):
                buy_trigger.update_trigger_price(amount)
            else:
                buy_trigger = None
        except ObjectDoesNotExist:
            buy_trigger = None

        return buy_trigger is not None
    
    def set_sell_trigger(self, symbol, amount):
        try:
            sell_trigger = SellTrigger.objects.get(stock_symbol=symbol,user_id=self.user_id)
            if(sell_trigger.cash_amount > 0):
                sell_trigger.update_trigger_price(amount)
        except ObjectDoesNotExist:
            sell_trigger = None

        return sell_trigger is not None

    def cancel_set_buy(self, symbol):
        try:
            buy_trigger = BuyTrigger.objects.get(stock_symbol=symbol,user_id=self.user_id)
            self.update_balance(buy_trigger.cash_amount)
            buy_trigger.cancel()
        except ObjectDoesNotExist:
            buy_trigger = None

        return buy_trigger is not None
    
    def cancel_set_sell(self, symbol):
        try:
            sell_trigger = SellTrigger.objects.get(stock_symbol=symbol,user_id=self.user_id)
            sell_trigger.cancel()
        except ObjectDoesNotExist:
            buy_trigger = None

        return buy_trigger is not None

    def update_balance(self, change):
        self.balance += change
        cache.set(self.user_id, self)
        self.save()

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
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    amount = models.PositiveIntegerField(default=0)

    def update_amount(self, change):
        self.amount += change
        self.save()


class SellTrigger(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    price = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_reserved_amount = models.PositiveIntegerField(default=0)
    committed = models.BooleanField(default=False)
    active = models.BooleanField(default=False)

    def check_validity(self, price):
        if(self.price <= price):
            user = User.get(user_id)
            sell = Sell.create(user_id=user.user_id, stock_symbol=self.stock_symbol,
                               cash_amount=cash_amount, stock_price=price)
            sell.commit(user)
            self.committed = True
            self.save()
    
    def update_cash_amount(self, amount):
        self.cash_amount += amount
        self.save()
    
    # TODO(isaacsahle): Chose to allow updating a set sell trigger
    # for the same stock. Since we need to reserve these stocks, 
    # a quote for the current price is necessarry. In the future,
    #  we might want to not allow updating these triggers to minimize
    #  quote server hits.
    def update_trigger_price(self, amount):
        user_stock = UserStock.objects.get(user_id=self.user_id, stock_symbol=self.stock_symbol)
        stock = Stock.quote(self.stock_symbol)
        stock_reserve_amount = amount // stock.price
        trigger_updated = False
        if stock_reserve_amount > 0 and (user_stock.amount + self.stock_reserved_amount) >= stock_reserve_amount:
            user_stock.update_amount(self.stock_reserved_amount - stock_reserve_amount)
            self.stock_reserved_amount = stock_reserve_amount
            self.price = amount
            self.active = True
            self.save()
            trigger_updated = True
        return trigger_updated

    def cancel(self):
        user_stock = UserStock.objects.get(user_id=self.user_id, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_reserved_amount)
        self.price = 0
        self.cash_amount = 0
        self.stock_reserved_amount = 0
        self.active = False
        self.save()



class BuyTrigger(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    price = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    committed = models.BooleanField(default=False)
    active = models.BooleanField(default=False)

    def check_validity(self, price, symbol):
        if(self.price >= price):
            user = User.get(user_id)
            buy = Buy.create(user_id=user.user_id, stock_symbol=self.stock_symbol,
                             cash_amount=cash_amount, stock_price=price)
            buy.commit(user)
            self.committed = True
            self.save()
    
    def update_cash_amount(self, amount):
        self.cash_amount += amount
        self.save()
   
    def update_trigger_price(self, amount):
        self.price = amount
        self.active = True
        self.save()
    
    def cancel(self):
        self.active = False
        self.price = 0
        self.cash_amount = 0
        self.save()

class Sell(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    intended_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    actual_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_sold_amount = models.PositiveIntegerField(default=0)
    sell_price = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)

    @classmethod
    def create(cls, stock_symbol, cash_amount, stock_price, user):
        sell = cls(user_id=user, stock_symbol=stock_symbol,
                   intended_cash_amount=cash_amount, sell_price=stock_price)
        sell.stock_sold_amount = Decimal(cash_amount)//Decimal(stock_price)
        sell.actual_cash_amount = Decimal(
            sell.stock_sold_amount)*(stock_price)
        sell.timestamp = time.time()
        user_stock = UserStock.objects.get(
            user_id=user, stock_symbol=stock_symbol)
        user_stock.update_amount(sell.stock_sold_amount*-1)
        return sell

    def commit(self, user):
        user.update_balance(self.actual_cash_amount)
        self.save()

    def cancel(self, user):
        user_stock = UserStock.objects.get(
            user_id=self.user_id, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_sold_amount)
        


class Buy(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    intended_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    actual_cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_bought_amount = models.PositiveIntegerField(default=0)
    purchase_price = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)

    @classmethod
    def create(cls, stock_symbol, cash_amount, stock_price, user):
        buy = cls(user_id=user, stock_symbol=stock_symbol,
                  purchase_price=stock_price)
        buy.stock_bought_amount = Decimal(cash_amount)//Decimal(stock_price)
        buy.actual_cash_amount = Decimal(buy.stock_bought_amount)*(stock_price)
        buy.timestamp = time.time()
        user.update_balance(buy.actual_cash_amount*-1)
        return buy

    def cancel(self, user):
        user.update_balance(self.actual_cash_amount)

    def commit(self, user):
        user_stock, created = UserStock.objects.get_or_create(
            user_id=user, stock_symbol=self.stock_symbol)
        user_stock.update_amount(self.stock_bought_amount)
        self.save()


def is_expired(previous_time):
    elapsed_time = time.time() - previous_time
    if(elapsed_time > 60):
        return True
    return False
