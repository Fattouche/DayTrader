
from .models import *
from django.core.cache import cache


def quote(symbol):
    stock = cache.get(symbol)
    if(stock is None):
        price = execute_quote_request(
            user_id=user_id, symbol=symbol)
        stock = Stock(symbol=symbol, price=price)
        cache.set(symbol, stock)
        django_rq.enqueue(stock.verify_triggers)
    return stock

# cailan implement this
def execute_quote_request(user_id, symbol):
    print("execute_quote_request not yet implemented")
    return 5
    
