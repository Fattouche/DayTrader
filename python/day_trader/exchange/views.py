import json
import threading

from django.shortcuts import render
from django.http import HttpResponse
import django_rq
from django.core.cache import cache
from decimal import Decimal
from django.http import JsonResponse

from .thread_local import get_current_logging_info, set_current_logging_info
from .audit_logging import AuditLogger
from .models.business_models import *


def store_logging_info(func):
    def wrapper(request):
        if request.method == 'POST':
            params = json.loads(request.body)
        elif request.method == 'GET':
            params = request.GET
        command = request.path[1:].upper()
        logging_info = {
            'transaction_num': params.get('transaction_num', -1),
            'server': 'BEAVER_1',  # TODO(cailan): get from environment var
            'command': command
        }
        set_current_logging_info(logging_info)
        return func(request)
    return wrapper


@store_logging_info
def index(request):
    return HttpResponse("Init view of stock exchange index")


@store_logging_info
def add(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    amount = Decimal(params.get('amount'))

    user = User.get(user_id)
    user.update_balance(amount)
    return JsonResponse({'action': 'add', 'balance': user.balance+amount}, status=200)


@store_logging_info
def quote(request):
    params = request.GET
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    stock = Stock.quote(symbol, user_id)
    return JsonResponse({'action': 'quote', 'stock': stock.symbol, 'price': stock.price}, status=200)


@store_logging_info
def buy(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = Decimal(params.get('amount'))

    user = User.get(user_id)
    err = user.perform_buy(symbol, amount)
    if err:
        return JsonResponse({'action': 'buy', 'error': err}, status=400)
    return JsonResponse({'action': 'buy', 'balance': user.balance, 'stock': symbol, 'price': user.buy_stack[-1].purchase_price, 'amount': amount, 'valid_duration': '60'}, status=200)


@store_logging_info
def commit_buy(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    user = User.get(user_id)
    buy = user.pop_from_buy_stack()
    if(not buy):
        return JsonResponse({'action': 'commit_buy', 'error': 'no buy currently exists'}, status=404)
    if(is_expired(buy.timestamp)):
        return JsonResponse({'action': 'commit_buy', 'stock': buy.stock_symbol, 'error': 'buy has expired, please re-buy in order to commit'}, status=408)
    buy.commit()
    return JsonResponse({'action': 'commit_buy', 'stock': buy.stock_symbol, 'amount_bought': buy.stock_bought_amount, 'balance': user.balance}, status=200)


@store_logging_info
def cancel_buy(request):
    params = json.loads(request.body)
    user = User.get(params.get('user_id'))
    user.cancel_buy()
    return JsonResponse({'action': 'cancel_buy', 'balance': user.balance}, status=200)


@store_logging_info
def sell(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = Decimal(params.get('amount'))
    user = User.get(user_id)
    err = user.perform_sell(symbol, amount)
    if err:
        return JsonResponse({'action': 'sell', 'error': err}, status=400)
    return JsonResponse({'action': 'sell', 'stock': symbol, 'price': user.sell_stack[-1].sell_price, 'amount': amount, 'valid_duration': '60'}, status=200)


@store_logging_info
def commit_sell(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    user = User.get(user_id)
    sell = user.pop_from_sell_stack()
    if(not sell):
        return JsonResponse({'action': 'commit_sell', 'error': 'no sell currently exists'}, status=404)
    if(is_expired(sell.timestamp)):
        return JsonResponse({'action': 'commit_sell', 'stock': sell.stock_symbol, 'error': 'sell has expired, please re-buy in order to commit'}, status=408)
    sell.commit(user)
    return JsonResponse({'action': 'commit_sell', 'stock': sell.stock_symbol, 'amount_sold': sell.stock_sold_amount, 'balance': user.balance}, status=200)


@store_logging_info
def cancel_sell(request):
    params = json.loads(request.body)
    user = User.get(params.get('user_id'))
    user.cancel_sell()
    return JsonResponse({'action': 'cancel_sell', 'balance': user.balance}, status=200)


@store_logging_info
def set_buy_amount(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = Decimal(params.get('amount'))
    user = User.get(user_id)
    err = user.set_buy_amount(symbol, amount)
    if err:
        return JsonResponse({'action': 'set_buy_amount', 'error': err}, status=412)
    return JsonResponse({'action': 'set_buy_amount', 'amount': amount}, status=200)


@store_logging_info
def set_sell_amount(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    amount = Decimal(params.get('amount'))
    user = User.get(user_id)
    err = user.set_sell_amount(symbol, amount)
    if err:
        return JsonResponse({'action': 'set_buy_amount', 'error': err}, status=412)
    return JsonResponse({'action': 'set_sell_amount', 'symbol': symbol, 'amount': amount}, status=200)


@store_logging_info
def set_buy_trigger(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    price = Decimal(params.get('amount'))
    user = User.get(user_id)
    err = user.set_buy_trigger(symbol, price)
    if err:
        return JsonResponse({'action': 'set_buy_trigger', 'error': err}, status=412)
    return JsonResponse({'action': 'set_buy_trigger', 'msg': 'buy trigger requires a set_buy_amount, please set one', 'symbol': symbol, 'price': price}, status=200)


@store_logging_info
def set_sell_trigger(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    price = Decimal(params.get('amount'))
    user = User.get(user_id)
    err = user.set_sell_trigger(symbol, price)
    if err:
        return JsonResponse({'action': 'set_sell_trigger', 'error': err}, status=412)
    return JsonResponse({'action': 'set_sell_trigger', 'msg': 'sell trigger requires a set_sell_amount, please set one', 'symbol': symbol, 'price': price}, status=200)


@store_logging_info
def cancel_set_buy(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    user = User.get(user_id)
    err = user.cancel_set_buy(symbol)
    if err:
        return JsonResponse({'action': 'cancel_set_buy', 'error': err}, status=404)
    return JsonResponse({'action': 'cancel_set_buy', 'symbol': symbol}, status=200)

@store_logging_info
def cancel_set_sell(request):
    params = json.loads(request.body)
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    user = User.get(user_id)
    err = user.cancel_set_sell(symbol)
    if err:
        return JsonResponse({'action': 'cancel_set_sell', 'error': err}, status=404)
    return JsonResponse({'action': 'cancel_set_sell', 'symbol': symbol}, status=200)


@store_logging_info
def dumplog(request):
    params = request.GET
    if 'filename' not in params:
        return JsonResponse({'error': 'filename is required'}, status=400)

    if 'user_id' not in params:
        AuditLogger.dump_system_logs(params['filename'])
    else:
        AuditLogger.dump_user_log(params['user_id'], params['filename'])

    return HttpResponse("System logs dumped to file {}".format(params['filename']))


@store_logging_info
def display_summary(request):
    return HttpResponse("Init view of stock exchange index")
