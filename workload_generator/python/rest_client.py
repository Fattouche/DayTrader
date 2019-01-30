import requests
import json


class RestClient:

    get_cmds = {"QUOTE": 1, "DUMPLOG": 1, "DISPLAY_SUMMARY": 1}

    symbol_cmds = {
        "QUOTE": 1,
        "BUY": 1,
        "SELL": 1,
        "SET_BUY_AMOUNT": 1,
        "CANCEL_SET_BUY": 1,
        "SET_BUY_TRIGGER": 1,
        "SET_SELL_AMOUNT": 1,
        "SET_SELL_TRIGGER": 1,
        "CANCEL_SET_SELL": 1
    }

    amount_cmds = {
        "ADD": 1,
        "BUY": 1,
        "SELL": 1,
        "SET_BUY_AMOUNT": 1,
        "SET_BUY_TRIGGER": 1,
        "SET_SELL_AMOUNT": 1,
        "SET_SELL_TRIGGER": 1,
    }

    def __init__(self, endpoint):
        self.endpoint = endpoint

    def send_cmd(self, user_id, cmd):

        cmd_type, req_data = self.build_request(user_id, *cmd.split(" "))
        url = self.build_url(self.endpoint, cmd_type)
        r = requests.get(url, params=req_data) if cmd_type in self.get_cmds else requests.post(
            url, data=json.dumps(req_data))

    def build_url(self, endpoint, cmd_type):
        url = "{}{}{}".format(endpoint, "/", cmd_type.lower())
        return url

    def build_request(self, user_id, *args):
        data = {
            "user_id": None,
            "symbol": None,
            "amount": None,
            "filename": None
        }

        cmd_type = args[0]

        data["user_id"] = user_id if user_id != "adminxxx" else None

        if cmd_type == "DUMPLOG":
            data["filename"] = args[2] if len(args) > 2 else args[1]

        if cmd_type in self.symbol_cmds:
            data["symbol"] = args[2] if len(args) > 2 else None

        if cmd_type in self.amount_cmds:
            if cmd_type == "ADD":
                data["amount"] = float(args[2]) if len(args) > 2 else None
            else:
                data["amount"] = float(args[3]) if len(args) > 3 else None

        return cmd_type, data
