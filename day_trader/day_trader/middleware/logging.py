import json
from exchange.audit_logging import AuditLogger


class LogRequestMiddleware(object):
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        command = request.path[1:].upper()
        log_input = {}
        if request.method == 'POST':
            params = json.loads(request.body)
        else:
            params = request.GET

        if 'user_id' in params:
            log_input['username'] = params['user_id']
        if 'symbol' in params:
            log_input['stock_symbol'] = params['symbol']
        if 'filename' in params:
            log_input['filename'] = params['filename']
        if 'funds' in params:
            log_input['funds'] = params['funds']
        logging_info = {
            'transaction_num': params.get('transaction_num', -1),
            'server': 'BEAVER_1'  # TODO(cailan): get from environment var
        }

        # TODO(cailan): deal with server_name
        AuditLogger.log_user_command(logging_info['server'],
                                     logging_info['transaction_num'], command, **(log_input))
        return self.get_response(request)
