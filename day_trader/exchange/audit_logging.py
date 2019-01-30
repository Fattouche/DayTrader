from lxml import etree
import exchange.models as Models

# TODO: Implement per-user dump
#
# TODO: write to DB tables using workers, rather than to CSV
#
# TODO: Consider buffering logs in memory, if DB I/O is a bottleneck
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
        def log_user_command(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds=None):
            base_log = Models.BaseLog(log_type='userCommand', 
                server=server_name, transaction_num=transaction_num)
            base_log.save()
            user_command_log = Models.UserCommandLog(
                command=command, username=username,
                stock_symbol=stock_symbol, filename=filename,
                funds=funds)
            user_command_log.base_log = base_log
            user_command_log.save()
            

        def log_quote_server_event(self, server_name, transaction_num,
                price, stock_symbol, username, quote_server_timestamp, 
                crypto_key):
            base_log = Models.BaseLog(log_type='quoteServer', 
                server=server_name, transaction_num=transaction_num)
            base_log.save()
            quote_server_log = Models.QuoteServerLog(
                price=price, stock_symbol=stock_symbol,
                username=username, 
                quote_server_time=quote_server_timestamp,
                crypto_key=crypto_key)
            quote_server_log.base_log = base_log
            quote_server_log.save()


        def log_account_transaction(self, server_name, transaction_num, action, 
                username, funds):
            base_log = Models.BaseLog(log_type='accountTransaction', 
                server=server_name, transaction_num=transaction_num)
            base_log.save()
            account_transaction_log = Models.AccountTransactionLog(
                action=action, username=username,
                funds=funds)
            account_transaction_log.base_log = base_log
            account_transaction_log.save()


        def log_system_event(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds=None):
            base_log = Models.BaseLog(log_type='systemEvent', 
                server=server_name, transaction_num=transaction_num)
            base_log.save()
            system_event_log = Models.SystemEventLog(
                command=command, username=username, stock_symbol=stock_symbol,
                filename=filename, funds=funds)
            system_event_log.base_log = base_log
            system_event_log.save()


        def log_error_event(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds=None, 
                error_message=''):
            base_log = Models.BaseLog(log_type='errorEvent', 
                server=server_name, transaction_num=transaction_num)
            base_log.save()
            error_event_log = Models.ErrorEventLog(
                command=command, username=username, stock_symbol=stock_symbol,
                filename=filename, funds=funds, error_message=error_message)
            error_event_log.base_log = base_log
            error_event_log.save()


        def log_debug_event(self, server_name, transaction_num, command, 
                username='', stock_symbol='', filename='', funds=None, 
                debug_message=''):
            base_log = Models.BaseLog(log_type='debugEvent', 
                server=server_name, transaction_num=transaction_num)
            base_log.save()
            debug_event_log = Models.DebugEventLog(
                command=command, username=username, stock_symbol=stock_symbol,
                filename=filename, funds=funds, debug_message=debug_message)
            debug_event_log.base_log = base_log
            debug_event_log.save()

        
        def dump_user_log(self, username, filename):
            pass


        def dump_system_logs(self, filename):
            with etree.xmlfile(filename) as xf:
                with xf.element('log'):
                    for user_command_log in Models.UserCommandLog.objects \
                                            .select_related('base_log').all():
                        base_log = user_command_log.base_log
                        AuditLogger._write_user_command_to_xml(xf, base_log, 
                            user_command_log)

                    for quote_server_log in Models.QuoteServerLog.objects \
                                            .select_related('base_log').all():
                        base_log = quote_server_log.base_log
                        AuditLogger._write_quote_server_event_to_xml(xf, 
                            base_log, quote_server_log)

                    for account_transaction_log in Models.AccountTransactionLog\
                                                    .objects \
                                                    .select_related('base_log')\
                                                    .all():
                        base_log = account_transaction_log.base_log
                        AuditLogger._write_account_transaction_to_xml(xf, 
                            base_log, account_transaction_log)

                    for system_event_log in Models.SystemEventLog.objects \
                                            .select_related('base_log').all():
                        base_log = system_event_log.base_log
                        AuditLogger._write_system_event_to_xml(xf, base_log, 
                            system_event_log)

                    for error_event_log in Models.ErrorEventLog.objects \
                                            .select_related('base_log').all():
                        base_log = error_event_log.base_log
                        AuditLogger._write_error_event_to_xml(xf, base_log, 
                            error_event_log)

                    for debug_event_log in Models.DebugEventLog.objects \
                                            .select_related('base_log').all():
                        base_log = debug_event_log.base_log
                        AuditLogger._write_debug_event_to_xml(xf, base_log, 
                            debug_event_log)


    @staticmethod
    def _write_xml_element(xf, name, text):
        element = etree.Element(name)
        element.text = text
        xf.write(element)


    @staticmethod
    def _write_user_command_to_xml(xf, base_log, user_command_log):
        with xf.element('userCommand'):
            AuditLogger._write_xml_element(xf, 'timestamp', 
                str(int(base_log.timestamp.timestamp() * 1000)))
            AuditLogger._write_xml_element(xf, 'server', base_log.server)
            AuditLogger._write_xml_element(xf, 'transactionNum', 
                str(base_log.transaction_num))
            AuditLogger._write_xml_element(xf, 'command', 
                user_command_log.command)
            if user_command_log.username:
                AuditLogger._write_xml_element(xf, 'username',
                    user_command_log.username)
            if user_command_log.stock_symbol:
                AuditLogger._write_xml_element(xf, 'stockSymbol', 
                    user_command_log.stock_symbol)
            if user_command_log.filename:
                AuditLogger._write_xml_element(xf, 'filename', 
                    user_command_log.filename)
            if user_command_log.funds:
                AuditLogger._write_xml_element(xf, 'funds', 
                    str(user_command_log.funds))


    @staticmethod
    def _write_quote_server_event_to_xml(xf, base_log, quote_server_log):
        with xf.element('quoteServer'):
            AuditLogger._write_xml_element(xf, 'timestamp', 
                str(int(base_log.timestamp.timestamp() * 1000)))
            AuditLogger._write_xml_element(xf, 'server', base_log.server)
            AuditLogger._write_xml_element(xf, 'transactionNum', 
                str(base_log.transaction_num))
            AuditLogger._write_xml_element(xf, 'price', 
                str(quote_server_log.price))
            AuditLogger._write_xml_element(xf, 'stockSymbol', 
                quote_server_log.stock_symbol)
            AuditLogger._write_xml_element(xf, 'username', 
                quote_server_log.username)
            AuditLogger._write_xml_element(xf, 'quoteServerTime', 
                str(quote_server_log.quote_server_time))
            AuditLogger._write_xml_element(xf, 'cryptokey',
                quote_server_log.crypto_key)


    @staticmethod
    def _write_account_transaction_to_xml(xf, base_log, 
            account_transaction_log):
        with xf.element('accountTransaction'):
            AuditLogger._write_xml_element(xf, 'timestamp', 
                str(int(base_log.timestamp.timestamp() * 1000)))
            AuditLogger._write_xml_element(xf, 'server', base_log.server)
            AuditLogger._write_xml_element(xf, 'transactionNum', 
                str(base_log.transaction_num))
            AuditLogger._write_xml_element(xf, 'action',
                account_transaction_log.action)
            AuditLogger._write_xml_element(xf, 'username',
                account_transaction_log.username)
            AuditLogger._write_xml_element(xf, 'funds',
                str(account_transaction_log.funds))


    @staticmethod
    def _write_system_event_to_xml(xf, base_log, system_event_log):
        with xf.element('systemEvent'):
            AuditLogger._write_xml_element(xf, 'timestamp', 
                str(int(base_log.timestamp.timestamp() * 1000)))
            AuditLogger._write_xml_element(xf, 'server', base_log.server)
            AuditLogger._write_xml_element(xf, 'transactionNum', 
                str(base_log.transaction_num))
            AuditLogger._write_xml_element(xf, 'command', 
                system_event_log.command)
            if system_event_log.username:
                AuditLogger._write_xml_element(xf, 'username',
                    system_event_log.username)
            if system_event_log.stock_symbol:
                AuditLogger._write_xml_element(xf, 'stockSymbol',
                    system_event_log.stock_symbol)
            if system_event_log.filename:
                AuditLogger._write_xml_element(xf, 'filename',
                    system_event_log.filename)
            if system_event_log.funds:
                AuditLogger._write_xml_element(xf, 'funds',
                    str(system_event_log.funds))


    @staticmethod
    def _write_error_event_to_xml(xf, base_log, error_event_log):
        with xf.element('errorEvent'):
            AuditLogger._write_xml_element(xf, 'timestamp', 
                str(int(base_log.timestamp.timestamp() * 1000)))
            AuditLogger._write_xml_element(xf, 'server', base_log.server)
            AuditLogger._write_xml_element(xf, 'transactionNum', 
                str(base_log.transaction_num))
            AuditLogger._write_xml_element(xf, 'command', 
                error_event_log.command)
            if error_event_log.username:
                AuditLogger._write_xml_element(xf, 'username',
                    error_event_log.username)
            if error_event_log.stock_symbol:
                AuditLogger._write_xml_element(xf, 'stockSymbol',
                    error_event_log.stock_symbol)
            if error_event_log.filename:
                AuditLogger._write_xml_element(xf, 'filename',
                    error_event_log.filename)
            if error_event_log.funds:
                AuditLogger._write_xml_element(xf, 'funds', 
                    str(error_event_log.funds))
            if error_event_log.error_message:
                AuditLogger._write_xml_element(xf, 'errorMessage', 
                    error_event_log.error_message)


    @staticmethod
    def _write_debug_event_to_xml(xf, base_log, debug_event_log):
        with xf.element('debugEvent'):
            AuditLogger._write_xml_element(xf, 'timestamp', 
                str(int(base_log.timestamp.timestamp() * 1000)))
            AuditLogger._write_xml_element(xf, 'server', base_log.server)
            AuditLogger._write_xml_element(xf, 'transactionNum', 
                str(base_log.transaction_num))
            AuditLogger._write_xml_element(xf, 'command', 
                debug_event_log.command)
            if debug_event_log.username:
                AuditLogger._write_xml_element(xf, 'username',
                    debug_event_log.username)
            if debug_event_log.stock_symbol:
                AuditLogger._write_xml_element(xf, 'stockSymbol',
                    debug_event_log.stock_symbol)
            if debug_event_log.filename:
                AuditLogger._write_xml_element(xf, 'filename',
                    debug_event_log.filename)
            if debug_event_log.funds:
                AuditLogger._write_xml_element(xf, 'funds', 
                    str(debug_event_log.funds))
            if debug_event_log.debug_message:
                AuditLogger._write_xml_element(xf, 'debugMessage', 
                    debug_event_log.debug_message)