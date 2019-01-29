import csv
from lxml import etree
from datetime import datetime

# Number of log entries to cache in memory before flushing to file
MAX_BUFFER_LEN = 1000

SYSTEM_OUTPUT_FILENAME = 'system_output.log'

# Encoding of transactions:
# Each transaction is represented as set of comma separated values, as 
# specified here:
# http://www.ece.uvic.ca/~seng468/ProjectWebSite/logfile.xsd
#
# Example: A BUY transaction, at time 12345, by user 'cool_dude', on server
# SERVER_X, with transaction number 100, with funds of 999, would be encoded as:
# accountTransaction,12345,SERVER_X,100,BUY,cool_dude,999
#
# TODO: Look into python coroutines for async IO: 
# https://lxml.de/api.html#incremental-xml-generation
#
# TODO: Implement per-user log files
#
# Note that AuditLogger is a singleton
class AuditLogger:
    _instance = None

    def __init__(self):
        if not AuditLogger._instance:
            AuditLogger._instance = AuditLogger._AuditLogger()
        else:
            raise Exception("AuditLogger is a singleton, only construct it "
                "once")
    

    @staticmethod
    def get_instance():
        if AuditLogger._instance is None:
            AuditLogger()

        return AuditLogger._instance


    class _AuditLogger:
        def __init__(self):
            # TODO: timestamp output filename, so that we don't overwrite
            # every time we run
            self.encoded_system_output = open(
                SYSTEM_OUTPUT_FILENAME, 'w', newline='')
            self.encoded_system_output_writer = csv.writer(
                self.encoded_system_output, delimiter=',')

            # Buffer some lines in memory, so that we aren't doing file IO for
            # every single transaction
            self.buffer = []


        def __del__(self):
            self.encoded_system_output.close()


        def log_user_command(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds=''):
            # TODO: log to user-specific file as well
            event = [
                server_name, str(transaction_num), command, username, 
                stock_symbol, filename, str(funds)
            ]
            self._log_event('userCommand', event)


        def log_quote_server_event(self, server_name, transaction_num,
                price, stock_symbol, username, quote_server_timestamp, 
                crypto_key):
            event = [
                server_name, str(transaction_num), str(price), stock_symbol,
                username, quote_server_timestamp, crypto_key
            ]
            self._log_event('quoteServer', event)


        def log_account_transaction(self, server_name, transaction_num, action, 
                username, funds):
            event = [
                server_name, str(transaction_num), action, username, 
                str(funds)
            ]
            self._log_event('accountTransaction', event)


        def log_system_event(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds=''):
            event = [
                server_name, str(transaction_num), command, username, 
                stock_symbol, filename, str(funds)
            ]
            self._log_event('systemEvent', event)


        def log_error_event(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds='', 
                error_message=''):
            event = [
                server_name, str(transaction_num), command, username, 
                stock_symbol, filename, str(funds), error_message
            ]
            self._log_event('errorEvent', event)


        def log_debug_event(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds='', 
                debug_message=''):
            event = [
                server_name, str(transaction_num), command, username, 
                stock_symbol, filename, str(funds), debug_message
            ]
            self._log_event('debugEvent', event)

        
        def dump_user_log(self, username, filename):
            pass


        def dump_system_logs(self, filename):
            self._dump_buffer()

            # Close encoded output file before we open it for reading
            self.encoded_system_output.close()

            with etree.xmlfile(filename) as xf:
                with xf.element('log'):
                    with open(SYSTEM_OUTPUT_FILENAME, newline='') as \
                            encoded_output:
                        reader = csv.reader(encoded_output)
                        for event_info in reader:
                            if event_info[0] == 'userCommand':
                                AuditLogger._decode_user_command_to_xml(
                                    xf, event_info)
                            elif event_info[0] == 'quoteServer':
                                AuditLogger._decode_quote_server_event_to_xml(
                                    xf, event_info)
                            elif event_info[0] == 'accountTransaction':
                                AuditLogger._decode_account_transaction_to_xml(
                                    xf, event_info)
                            elif event_info[0] == 'systemEvent':
                                AuditLogger._decode_system_event_to_xml(
                                    xf, event_info)
                            elif event_info[0] == 'errorEvent':
                                AuditLogger._decode_error_event_to_xml(
                                    xf, event_info)
                            elif event_info[0] == 'debugEvent':
                                AuditLogger._decode_debug_event_to_xml(
                                    xf, event_info)

            
            # Re-open for writing once we're done
            self.encoded_system_output = open(SYSTEM_OUTPUT_FILENAME, 'a')
            self.encoded_system_output_writer = csv.writer(
                self.encoded_system_output)


        def _dump_buffer(self):
            self.encoded_system_output_writer.writerows(self.buffer)
            self.buffer.clear()


        def _log_event(self, event_type, event):
            curr_timestamp = int(datetime.utcnow().timestamp() * 1000)
            self.buffer.append(
                [event_type, str(curr_timestamp)] + 
                event
            )

            if len(self.buffer) >= MAX_BUFFER_LEN:
                self._dump_buffer()  


    @staticmethod
    def _write_xml_element(xf, event_info, index, name):
        element = etree.Element(name)
        element.text = event_info[index]
        xf.write(element)


    @staticmethod
    def _decode_user_command_to_xml(xf, event_info):
        with xf.element('userCommand'):
            AuditLogger._write_xml_element(xf, event_info, 1, 'timestamp')
            AuditLogger._write_xml_element(xf, event_info, 2, 'server')
            AuditLogger._write_xml_element(xf, event_info, 3, 'transactionNum')
            AuditLogger._write_xml_element(xf, event_info, 4, 'command')
            if event_info[5]:
                AuditLogger._write_xml_element(xf, event_info, 5, 'username')
            if event_info[6]:
                AuditLogger._write_xml_element(xf, event_info, 6, 'stockSymbol')
            if event_info[7]:
                AuditLogger._write_xml_element(xf, event_info, 7, 'filename')
            if event_info[8]:
                AuditLogger._write_xml_element(xf, event_info, 8, 'funds')


    @staticmethod
    def _decode_quote_server_event_to_xml(xf, event_info):
        with xf.element('quoteServer'):
            AuditLogger._write_xml_element(xf, event_info, 1, 'timestamp')
            AuditLogger._write_xml_element(xf, event_info, 2, 'server')
            AuditLogger._write_xml_element(xf, event_info, 3, 'transactionNum')
            AuditLogger._write_xml_element(xf, event_info, 4, 'price')
            AuditLogger._write_xml_element(xf, event_info, 5, 'stockSymbol')
            AuditLogger._write_xml_element(xf, event_info, 6, 'username')
            AuditLogger._write_xml_element(xf, event_info, 7, 'quoteServerTime')
            AuditLogger._write_xml_element(xf, event_info, 8, 'cryptokey')


    @staticmethod
    def _decode_account_transaction_to_xml(xf, event_info):
        with xf.element('accountTransaction'):
            AuditLogger._write_xml_element(xf, event_info, 1, 'timestamp')
            AuditLogger._write_xml_element(xf, event_info, 2, 'server')
            AuditLogger._write_xml_element(xf, event_info, 3, 'transactionNum')
            AuditLogger._write_xml_element(xf, event_info, 4, 'action')
            AuditLogger._write_xml_element(xf, event_info, 5, 'username')
            AuditLogger._write_xml_element(xf, event_info, 6, 'funds')


    @staticmethod
    def _decode_system_event_to_xml(xf, event_info):
        with xf.element('systemEvent'):
            AuditLogger._write_xml_element(xf, event_info, 1, 'timestamp')
            AuditLogger._write_xml_element(xf, event_info, 2, 'server')
            AuditLogger._write_xml_element(xf, event_info, 3, 'transactionNum')
            AuditLogger._write_xml_element(xf, event_info, 4, 'command')
            if event_info[5]:
                AuditLogger._write_xml_element(xf, event_info, 5, 'username')
            if event_info[6]:
                AuditLogger._write_xml_element(xf, event_info, 6, 'stockSymbol')
            if event_info[7]:
                AuditLogger._write_xml_element(xf, event_info, 7, 'filename')
            if event_info[8]:
                AuditLogger._write_xml_element(xf, event_info, 8, 'funds')


    @staticmethod
    def _decode_error_event_to_xml(xf, event_info):
        with xf.element('errorEvent'):
            AuditLogger._write_xml_element(xf, event_info, 1, 'timestamp')
            AuditLogger._write_xml_element(xf, event_info, 2, 'server')
            AuditLogger._write_xml_element(xf, event_info, 3, 'transactionNum')
            AuditLogger._write_xml_element(xf, event_info, 4, 'command')
            if event_info[5]:
                AuditLogger._write_xml_element(xf, event_info, 5, 'username')
            if event_info[6]:
                AuditLogger._write_xml_element(xf, event_info, 6, 'stockSymbol')
            if event_info[7]:
                AuditLogger._write_xml_element(xf, event_info, 7, 'filename')
            if event_info[8]:
                AuditLogger._write_xml_element(xf, event_info, 8, 'funds')
            if event_info[9]:
                AuditLogger._write_xml_element(
                    xf, event_info, 9, 'errorMessage')


    @staticmethod
    def _decode_debug_event_to_xml(xf, event_info):
        with xf.element('debugEvent'):
            AuditLogger._write_xml_element(xf, event_info, 1, 'timestamp')
            AuditLogger._write_xml_element(xf, event_info, 2, 'server')
            AuditLogger._write_xml_element(xf, event_info, 3, 'transactionNum')
            AuditLogger._write_xml_element(xf, event_info, 4, 'command')
            if event_info[5]:
                AuditLogger._write_xml_element(xf, event_info, 5, 'username')
            if event_info[6]:
                AuditLogger._write_xml_element(xf, event_info, 6, 'stockSymbol')
            if event_info[7]:
                AuditLogger._write_xml_element(xf, event_info, 7, 'filename')
            if event_info[8]:
                AuditLogger._write_xml_element(xf, event_info, 8, 'funds')
            if event_info[9]:
                AuditLogger._write_xml_element(
                    xf, event_info, 9, 'debugMessage')