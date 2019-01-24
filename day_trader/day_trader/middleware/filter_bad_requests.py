class FilterBadRequestsMiddleware(object):
    def __init__(self, get_response):
        self.get_response = get_response
    
    def __call__(self, request):        
         # Stuff to do before a request. 
         response = self.get_response(request)
         # Stuff to do after a request.
         return response