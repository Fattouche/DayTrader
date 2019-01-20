from django.db import models


class Stock(models.Model):
    symbol = models.CharField(max_length=3, primary_key=True)
    price = models.PositiveIntegerField(default=0)


class User(models.Model):
    user_id = models.CharField(max_length=100, primary_key=True)
    balance = models.PositiveIntegerField(default=0)
    stocks = models.ManyToManyField(Stock)


class Trigger(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    price = models.PositiveIntegerField(default=0)
    sell_amount = models.PositiveIntegerField(default=0)
    trigger_type = models.TextField(default="buy")


class Sell(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    cash_amount = models.PositiveIntegerField(default=0)
    stock_sold_amount = models.PositiveIntegerField(default=0)
    commited = models.BooleanField(default=False)


class Buy(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    stock_symbol = models.ForeignKey(Stock, on_delete=models.CASCADE)
    cash_amount = models.PositiveIntegerField(default=0)
    commited = models.BooleanField(default=False)
