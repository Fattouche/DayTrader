from django.shortcuts import render

from django.http import HttpResponse


def index(request):
    return HttpResponse("Init view of stock exchange index")


def add(request):
    return HttpResponse("Init view of stock exchange index")


def quote(request):
    return HttpResponse("Init view of stock exchange index")


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


def set_sell_trigger(request):
    return HttpResponse("Init view of stock exchange index")


def set_buy_trigger(request):
    return HttpResponse("Init view of stock exchange index")


def dumplog(request):
    return HttpResponse("Init view of stock exchange index")


def display_summary(request):
    return HttpResponse("Init view of stock exchange index")
