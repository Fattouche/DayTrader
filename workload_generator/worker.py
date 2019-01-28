import threading
import rest_client

def target(user_id, cmd_list):
    client = rest_client.RestClient("http://localhost")
    for cmd in cmd_list:
        client.send_cmd(user_id, cmd)  
        