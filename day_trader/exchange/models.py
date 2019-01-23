from django.db import models


class Stock(models.Model):
    symbol = models.CharField(max_length=3, primary_key=True)
    price = models.FloatField(default=0)

    def verify_triggers(self, increased):
        # Only need to check sell
        if(increased):
            sell_trigger = SellTrigger.objects.filter(
                stock_symbol=self.symbol)
            for trigger in sell_trigger:
                if(trigger.price <= self.price):
                    user = User.objects.get(user_id=trigger.user_id)
                    user.update_balance(self.price*trigger.stock_amount)
                    trigger.committed = True
                    trigger.save()
        else:
            buy_trigger = BuyTrigger.objects.filter(
                stock_symbol=self.symbol, trigger_type="buy")
            for trigger in buy_trigger:
                if(trigger.price >= self.price):
                    user_stock = UserStock.objects.get(
                        user_id=trigger.user_id, stock_symbol=self.symbol)
                    user_stock.update_amount(
                        trigger.cash_amount//self.price)
                    trigger.committed = True
                    trigger.save()


class User(models.Model):
    user_id = models.CharField(max_length=100, primary_key=True)
    balance = models.FloatField(default=0)
    stocks = models.ManyToManyField(Stock, through='UserStock')
    password = models.CharField(max_length=50)

    def update_balance(self, change):
        self.balance += change
        self.save()


class UserStock(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    amount = models.PositiveIntegerField(default=0)

    def update_amount(self, change):
        self.amount += change
        self.save()


class SellTrigger(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    price = models.FloatField(default=0)
    stock_amount = models.FloatField(default=0)
    committed = models.BooleanField(default=False)


class BuyTrigger(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    price = models.FloatField(default=0)
    cash_amount = models.FloatField(default=0)
    committed = models.BooleanField(default=False)


class Sell(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    cash_amount = models.FloatField(default=0)
    stock_sold_amount = models.PositiveIntegerField(default=0)
    committed = models.BooleanField(default=False)


class Buy(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    cash_amount = models.FloatField(default=0)
    committed = models.BooleanField(default=False)
