import threading

def target(user_id, cmd_list):
    print(threading.current_thread().getName() , user_id, len(cmd_list))