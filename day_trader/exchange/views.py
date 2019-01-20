from django.shortcuts import render

from django.http import HttpResponse


def index(request):
    return HttpResponse("Init view of stock exchange index")
