from django.urls import path

from . import views

# Can we have it dynamically choose the function based off the url path?
urlpatterns = [
    path('', views.index, name='index'),
    path('add', views.add, name='add'),
    path('quote', views.quote, name='quote'),
    path('buy', views.buy, name='buy'),
    path('commit_buy', views.commit_buy, name='commit_buy'),
    path('cancel_buy', views.cancel_buy, name='cancel_buy'),
    path('sell', views.sell, name='sell'),
    path('commit_sell', views.commit_sell, name='commit_sell'),
    path('cancel_sell', views.cancel_sell, name='cancel_sell'),
    path('cancel_buy', views.cancel_sell, name='cancel_buy'),
    path('set_buy_amount', views.set_buy_amount, name='set_buy_amount'),
    path('set_sell_amount', views.set_sell_amount, name='set_sell_amount'),
    path('set_sell_trigger', views.set_sell_trigger, name='set_sell_trigger'),
    path('cancel_set_sell', views.cancel_set_sell, name='cancel_set_sell'),
    path('cancel_set_buy', views.cancel_set_buy, name='cancel_set_buy'),
    path('set_buy_trigger', views.set_buy_trigger, name='set_buy_trigger'),
    path('dumplog', views.dumplog, name='dumplog'),
    path('display_summary', views.display_summary, name='display_summary'),
]
