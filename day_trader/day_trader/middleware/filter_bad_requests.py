from django.http import HttpResponseBadRequest
from decimal import Decimal

class FilterBadRequestsMiddleware(object):
    def __init__(self, get_response):
        self.get_response = get_response
    
    def __call__(self, request):
        queryDict = request.POST if request.method == 'POST' else request.GET
        if 'amount' in queryDict and Decimal(queryDict['amount']) < 0:
            return HttpResponseBadRequest("Amount should be positive")
        if 'symbol' in queryDict and (not queryDict['symbol'] or len(queryDict['symbol']) > 3 or not queryDict['symbol'].isalpha()):
            return HttpResponseBadRequest("Symbol should be 1-3 alphabetic characters")

        return self.get_response(request)