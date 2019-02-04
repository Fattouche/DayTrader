import threading
_thread_locals = threading.local()

def get_current_logging_info():
    return getattr(_thread_locals, 'logging_info', None)

def set_current_logging_info(logging_info):
    _thread_locals.logging_info = logging_info