import threading
import rest_client


def target(user_id, cmd_list):

    for cmd in cmd_list:
        rest_client.send_cmd(cmd)    
    


