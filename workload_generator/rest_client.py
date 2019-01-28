import requests
import json 

get_cmds = {"QUOTE":1, "DUMPLOG":1, "DISPLAY_SUMMARY":1}

symbol_cmds = {
    "QUOTE":1,
    "BUY":1,
    "SELL":1, 
    "SET_BUY_AMOUNT":1, 
    "CANCEL_SET_BUY":1, 
    "SET_BUY_TRIGGER":1, 
    "SET_SELL_AMOUNT":1, 
    "SET_SELL_TRIGGER":1, 
    "CANCEL_SET_SELL":1
    }

amount_cmds = {
   "ADD":1,
   "BUY":1,
   "SELL":1,
   "SET_BUY_AMOUNT":1,
   "SET_BUY_TRIGGER":1,
   "SET_SELL_AMOUNT":1,
   "SET_SELL_TRIGGER":1,
}

endpoint = "http://localhost"

def send_cmd(cmd):

    cmd_type, req_data = build_request(*cmd.split(" "))
    url = build_url(endpoint, cmd_type)

    if cmd_type in get_cmds:
        r = requests.get(url, data=json.dumps(req_data))
    else:
        r = requests.post(url, data=json.dumps(req_data))

def build_url(endpoint, cmd_type):
    url = endpoint + "/" + cmd_type.lower()
    return url

def build_request(*args):
    data = 
        {
        "user_id": None, 
        "symbol": None,
        "amount": None,
        "filename": None
        }

    cmd_type =  args[0]

    if cmd_type!= "DUMPLOG":
        data["user_id"] = args[1]
    else:
        data["user_id"] = args[1] if len(args) > 2 else None
        data["filename"] = args[2] if len(args) > 2 else args[1]

    if cmd_type in symbol_cmds:
        data["symbol"] = args[2] if len(args) > 2 else None
    
    if cmd_type in amount_cmds:
        if cmd_type == "ADD":
            data["amount"] = float(args[2]) if len(args) > 2 else None
        else:
            data["amount"] = float(args[3]) if len(args) > 3 else None

    return cmd_type, data