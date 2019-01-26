from django.shortcuts import render

from django.http import HttpResponse

import django_rq

from .models import *
from django.core.cache import cache
from decimal import Decimal
from django.http import JsonResponse


def index(request):
    return HttpResponse("Init view of stock exchange index")

def add(request):
    params = request.POST
    user_id = params.get('user_id')
    amount = Decimal(params.get('amount'))
    user = User.get(user_id)
    user.update_balance(amount)
    return JsonResponse({'action':'add', 'balance': user.balance}, status=200)

def quote(request):
    params = request.GET
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    stock = Stock.quote(symbol)
    return JsonResponse({'action':'quote', 'stock': stock.symbol, 'price': stock.price}, status=200)


def buy(request):
    params = request.POST
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = params.get('amount')
    user = User.get(user_id)
    user.perform_buy(symbol, amount)
    return JsonResponse({'action':'buy', 'stock': stock.symbol, 'amount': amount, 'valid_duration':'60'}, status=200)


def commit_buy(request):
    params = request.POST
    user_id = params.get('user_id')
    user = User.get(user_id)
    buy = user.buy_stack.pop()
    if(is_expired(buy.timestamp)):
         return JsonResponse({'action':'commit_buy', 'stock': buy.stock_symbol, 'error':'buy has expired, please re-buy in order to commit'}, status=408)
    buy.commit(user)
    return JsonResponse({'action':'commit_buy', 'stock': buy.stock_symbol, 'amount':buy.stock_bought_amount, 'balance':user.balance}, status=200)


def cancel_buy(request):
    return HttpResponse("Init view of stock exchange index")


def sell(request):
    params = request.POST
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = params.get('amount')
    user = User.get(user_id)
    user.perform_sell(symbol, amount)
    return JsonResponse({'action':'sell', 'stock': stock.symbol, 'amount': amount, 'valid_duration':'60'}, status=200)


def commit_sell(request):
    params = request.POST
    user_id = params.get('user_id')
    user = User.get(user_id)
    sell = user.buy_stack.pop()
    if(is_expired(sell.timestamp)):
         return JsonResponse({'action':'commit_sell', 'stock': sell.stock_symbol, 'error':'sell has expired, please re-buy in order to commit'}, status=408)
    sell.commit(user)
    return JsonResponse({'action':'commit_sell', 'stock': sell.stock_symbol, 'amount':sell.stock_sold_amount, 'balance':user.balance}, status=200)


def cancel_sell(request):
    return HttpResponse("Init view of stock exchange index")


def set_buy_amount(request):
    return HttpResponse("Init view of stock exchange index")


def set_sell_amount(request):
    return HttpResponse("Init view of stock exchange index")


def set_sell_trigger(request):
    return HttpResponse("Init view of stock exchange index")


def cancel_set_sell(request):
    return HttpResponse("Init view of stock exchange index")


def set_buy_trigger(request):
    return HttpResponse("Init view of stock exchange index")


def dumplog(request):
    return HttpResponse("Init view of stock exchange index")


def display_summary(request):
    return HttpResponse("Init view of stock exchange index")
