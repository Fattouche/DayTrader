from django.shortcuts import render

from django.http import HttpResponse

import django_rq

from .models import *
from .utils import *
from django.core.cache import cache
from decimal import Decimal
from django.http import JsonResponse


def index(request):
    return HttpResponse("Init view of stock exchange index")


def add(request):
    params = request.POST
    user_id = params.get('user_id')
    amount = params.get('amount')
    user = User.objects.get(user_id=user_id)
    user.balance += Decimal(amount)
    user.save()
    return JsonResponse({'balance': user.balance}, status=200)


def quote(request):
    params = request.GET
    user_id = params.get('user_id')
    symbol = params.get('symbol')
    stock = quote(symbol)
    return JsonResponse({'stock': stock.symbol, 'price': stock.price}, status=200)


def buy(request):
    return HttpResponse("Init view of stock exchange index")


def commit_buy(request):
    return HttpResponse("Init view of stock exchange index")


def cancel_buy(request):
    return HttpResponse("Init view of stock exchange index")


def sell(request):
    return HttpResponse("Init view of stock exchange index")


def commit_sell(request):
    return HttpResponse("Init view of stock exchange index")


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
