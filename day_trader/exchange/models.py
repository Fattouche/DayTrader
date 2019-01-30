from django.db import models

class Stack:
    def __init__(self, user_id):
        self.user_id = user_id
        self.buy_stack = []
        self.sell_stack = []

class Stock:
    def __init__(self, symbol, price):
        self.symbol = symbol
        self.price = price

    def check_sell_trigger(self):
        sell_trigger = SellTrigger.objects.filter(
            stock_symbol=self.symbol, committed=False)
        for trigger in sell_trigger:
            trigger.check_validity(self.price)

    def check_buy_trigger(self):
        buy_trigger = BuyTrigger.objects.filter(
            stock_symbol=self.symbol, committed=False)
        for trigger in buy_trigger:
            trigger.check_validity(self.price, self.symbol)

    def verify_triggers(self):
        self.check_sell_trigger()
        self.check_buy_trigger()


class User(models.Model):
    user_id = models.CharField(max_length=100, primary_key=True)
    balance = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    password = models.CharField(max_length=50)

    def update_balance(self, change):
        self.balance += change
        self.save()


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
    stock_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    committed = models.BooleanField(default=False)

    def check_validity(self, price):
        if(self.price <= price):
            user = User.objects.get(user_id=trigger.user_id)
            user.update_balance(price*self.stock_amount)
            self.committed = True
            self.save()


class BuyTrigger(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    price = models.DecimalField(max_digits=65, decimal_places=2, default=0)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    committed = models.BooleanField(default=False)

    def check_validity(self, price, symbol):
        if(self.price >= price):
            user_stock = UserStock.objects.get(
                user_id=self.user_id, stock_symbol=symbol)
            user_stock.update_amount(cash_amount//self.price)
            self.committed = True
            self.save()


class Sell(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)
    stock_sold_amount = models.PositiveIntegerField(default=0)


class Buy(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.CharField(max_length=3, primary_key=True)
    cash_amount = models.DecimalField(
        max_digits=65, decimal_places=2, default=0)


class BaseLog(models.Model):
    log_type = models.CharField(max_length=32)
    timestamp = models.DateTimeField(auto_now_add=True)
    server = models.CharField(max_length=8)
    transaction_num = models.PositiveIntegerField()


class UserCommandLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.IntegerField(null=True)


class QuoteServerLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    price = models.PositiveIntegerField()
    stock_symbol = models.CharField(max_length=3)
    username = models.CharField(max_length=128)
    quote_server_time = models.PositiveIntegerField()
    crypto_key = models.CharField(max_length=256)  # TODO: what is crypto key length?


class AccountTransactionLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    action = models.CharField(max_length=32)
    username = models.CharField(max_length=64)
    funds = models.IntegerField()


class SystemEventLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.IntegerField(null=True)


class ErrorEventLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.IntegerField(null=True)
    error_message = models.CharField(max_length=512)


class DebugEventLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.IntegerField(null=True)
    debug_message = models.CharField(max_length=512)
