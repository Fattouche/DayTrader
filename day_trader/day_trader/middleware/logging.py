from exchange.audit_logging import AuditLogger

class LogRequestMiddleware(object):
    def __init__(self, get_response):
        self.get_response = get_response
        self.logger = AuditLogger.get_instance()

    def __call__(self, request):
        params = json.loads(request.body)
        # TODO(cailan): deal with server_name
        # TODO(cailan): deal with transaction_num
        log_input = {}
        if 'user_id' in params:
            log_input['username'] = params['user_id']
        if 'symbol' in params:
            log_input['stock_symbol'] = params['symbol']
        if 'filename' in params:
            log_input['filename'] = params['filename']
        if 'funds' in params:
            log_input['funds'] = params['funds']

        log_user_command('BEAVER_1', 1, params['command'], **(log_input))
