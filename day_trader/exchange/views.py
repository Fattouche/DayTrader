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
    return JsonResponse({'action': 'add', 'balance': user.balance}, status=200)


def quote(request):
    params = request.GET
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    stock = Stock.quote(symbol, user_id)
    return JsonResponse({'action': 'quote', 'stock': stock.symbol, 'price': stock.price}, status=200)


def buy(request):
    params = request.POST
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = params.get('amount')
    user = User.get(user_id)
    err = user.perform_buy(symbol, amount)
    if err is not None:
        return JsonResponse({'action': 'buy', 'error': err}, status=400)
    return JsonResponse({'action': 'buy', 'balance': user.balance, 'stock': symbol, 'price': user.buy_stack[-1].purchase_price, 'amount': amount, 'valid_duration': '60'}, status=200)


def commit_buy(request):
    params = request.POST
    user_id = params.get('user_id')
    user = User.get(user_id)
    buy = user.pop_from_buy_stack()
    if(buy is None):
        return JsonResponse({'action': 'commit_buy', 'error': 'no buy currently exists'}, status=404)
    if(is_expired(buy.timestamp)):
        return JsonResponse({'action': 'commit_buy', 'stock': buy.stock_symbol, 'error': 'buy has expired, please re-buy in order to commit'}, status=408)
    buy.commit(user)
    return JsonResponse({'action': 'commit_buy', 'stock': buy.stock_symbol, 'amount_bought': buy.stock_bought_amount, 'balance': user.balance}, status=200)


def cancel_buy(request):
    return HttpResponse("Init view of stock exchange index")


def sell(request):
    params = request.POST
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = params.get('amount')
    user = User.get(user_id)
    sell = user.pop_from_sell_stack()
    err = user.perform_sell(symbol, amount)
    if err is not None:
        return JsonResponse({'action': 'sell', 'error': err}, status=400)
    return JsonResponse({'action': 'sell', 'stock': symbol, 'price': user.sell_stack[-1].sell_price, 'amount': amount, 'valid_duration': '60'}, status=200)


def commit_sell(request):
    params = request.POST
    user_id = params.get('user_id')
    user = User.get(user_id)
    sell = user.pop_from_sell_stack()
    if(sell is None):
        return JsonResponse({'action': 'commit_sell', 'error': 'no sell currently exists'}, status=404)
    if(is_expired(sell.timestamp)):
        return JsonResponse({'action': 'commit_sell', 'stock': sell.stock_symbol, 'error': 'sell has expired, please re-buy in order to commit'}, status=408)
    sell.commit(user)
    return JsonResponse({'action': 'commit_sell', 'stock': sell.stock_symbol, 'amount_sold': sell.stock_sold_amount, 'balance': user.balance}, status=200)


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
