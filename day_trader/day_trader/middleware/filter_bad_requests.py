from django.http import JsonResponse
from django.urls import reverse
from decimal import Decimal
import json


class FilterBadRequestsMiddleware(object):
    def __init__(self, get_response):
        self.get_response = get_response
        self.index = reverse('index')
        self.admin_index = reverse('admin:index')
        self.paths_requiring_symbol_and_amount = {
            reverse('buy'),
            reverse('sell'),
            reverse('set_buy_amount'),
            reverse('set_buy_trigger'),
            reverse('set_sell_amount'),
            reverse('set_sell_trigger'),
        }
        self.add_path = reverse('add')
        self.quote_path = reverse('quote')
        self.dumplog_path = reverse('dumplog')

    def __call__(self, request):
        query_dict = json.loads(
            request.body) if request.method == 'POST' else request.GET

        if 'user_id' not in query_dict and request.path != self.index and not request.path.startswith(self.admin_index):
            return JsonResponse({'error': 'User id is required'}, status=400)
        hasSymbol = 'symbol' in query_dict
        hasAmount = 'amount' in query_dict
        if not (hasSymbol and hasAmount) and request.path in self.paths_requiring_symbol_and_amount:
            return JsonResponse({'error': 'Symbol and amount is required'}, status=400)
        if not hasSymbol and request.path == self.quote_path:
            return JsonResponse({'error': 'Symbol is required'}, status=400)
        if not hasAmount and request.path == self.add_path:
            return JsonResponse({'error': 'Amount is required'}, status=400)
        if hasAmount and Decimal(query_dict['amount']) < 0:
            return JsonResponse({'error': 'Amount should be positive'}, status=400)
        if hasSymbol and (not query_dict['symbol'] or len(query_dict['symbol']) > 3 or not query_dict['symbol'].isalpha()):
            return JsonResponse({'error': 'Symbol should be 1-3 alphabetic characters'}, status=400)
        if 'filename' not in query_dict and request.path == self.dumplog_path:
            return JsonResponse({'error': 'Filename is required'}, status=400)

        return self.get_response(request)
