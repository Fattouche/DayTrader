from django.http import HttpResponseBadRequest
from django.urls import reverse
from decimal import Decimal

class FilterBadRequestsMiddleware(object):
    def __init__(self, get_response):
        self.get_response = get_response
    
    def __call__(self, request):
        query_dict = request.POST if request.method == 'POST' else request.GET
        
        if 'amount' in query_dict and Decimal(query_dict['amount']) < 0:
            return HttpResponseBadRequest('Amount should be positive')
        if 'symbol' in query_dict and (not query_dict['symbol'] or len(query_dict['symbol']) > 3 or not query_dict['symbol'].isalpha()):
            return HttpResponseBadRequest('Symbol should be 1-3 alphabetic characters') 
        if 'user_id' not in query_dict and request.path != reverse('index') and not request.path.startswith(reverse('admin:index')):
            return HttpResponseBadRequest('User id is required')
        
        return self.get_response(request)