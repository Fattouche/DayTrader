import threading
_thread_locals = threading.local()

def get_current_request():
    return getattr(_thread_locals, 'request', None)

class ThreadLocals(object):
    """
    Middleware that gets various objects from the
    request object and saves them in thread local storage.
    """
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        self.process_request(request)
        return self.get_response(request)

    def process_request(self, request):
        print('processed request')
        _thread_locals.request = request