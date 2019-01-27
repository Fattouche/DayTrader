from django.test import TestCase
from .audit_logging import AuditLogger
import xml.etree.ElementTree as ET

class AuditLoggingTestCase(TestCase):
    def setUp(self):
        self.audit_logger = AuditLogger.get_instance()

    def test_logging_and_system_dumplog(self):
        self.audit_logger.log_user_command('SERVER_1', 1, 'BUY', 
            username='cailan', stock_symbol='HGU', filename='blah.txt',
            funds=1000000)
        self.audit_logger.log_quote_server_event('SERVER_1', 1, 25, 'HGU',
            'cailan', '12345', 'key_15')
        self.audit_logger.log_account_transaction('SERVER_1', 1, 'BUY',
            'cailan', 1000000)
        self.audit_logger.log_system_event('SERVER_1', 1, 'BUY',
            username='cailan', stock_symbol='HGU')
        self.audit_logger.log_error_event('SERVER_1', 1, 'BUY',
            username='cailan', stock_symbol='HGU', error_message='Oops')
        self.audit_logger.log_debug_event('SERVER_1', 1, 'BUY',
            username='cailan', stock_symbol='HGU', debug_message='Complete')

        self.audit_logger.dump_system_logs('test_output.xml')

        tree = ET.parse('test_output.xml')
        root = tree.getroot()

        self.assertEqual(root.tag, 'log')

        self.assertEqual(root[0].tag, 'userCommand')
        self.assertEqual(root[0][0].tag, 'timestamp')
        self.assertEqual(root[0][1].tag, 'server')
        self.assertEqual(root[0][1].text, 'SERVER_1')
        self.assertEqual(root[0][2].tag, 'transactionNum')
        self.assertEqual(root[0][2].text, '1')
        self.assertEqual(root[0][3].tag, 'command')
        self.assertEqual(root[0][3].text, 'BUY')
        self.assertEqual(root[0][4].tag, 'username')
        self.assertEqual(root[0][4].text, 'cailan')
        self.assertEqual(root[0][5].tag, 'stockSymbol')
        self.assertEqual(root[0][5].text, 'HGU')
        self.assertEqual(root[0][6].tag, 'filename')
        self.assertEqual(root[0][6].text, 'blah.txt')
        self.assertEqual(root[0][7].tag, 'funds')
        self.assertEqual(root[0][7].text, '1000000')

        self.assertEqual(root[1].tag, 'quoteServer')
        self.assertEqual(root[1][0].tag, 'timestamp')
        self.assertEqual(root[1][1].tag, 'server')
        self.assertEqual(root[1][1].text, 'SERVER_1')
        self.assertEqual(root[1][2].tag, 'transactionNum')
        self.assertEqual(root[1][2].text, '1')
        self.assertEqual(root[1][3].tag, 'price')
        self.assertEqual(root[1][3].text, '25')
        self.assertEqual(root[1][4].tag, 'stockSymbol')
        self.assertEqual(root[1][4].text, 'HGU')
        self.assertEqual(root[1][5].tag, 'username')
        self.assertEqual(root[1][5].text, 'cailan')
        self.assertEqual(root[1][6].tag, 'quoteServerTime')
        self.assertEqual(root[1][6].text, '12345')
        self.assertEqual(root[1][7].tag, 'cryptokey')
        self.assertEqual(root[1][7].text, 'key_15')

        self.assertEqual(root[2].tag, 'accountTransaction')
        self.assertEqual(root[2][0].tag, 'timestamp')
        self.assertEqual(root[2][1].tag, 'server')
        self.assertEqual(root[2][1].text, 'SERVER_1')
        self.assertEqual(root[2][2].tag, 'transactionNum')
        self.assertEqual(root[2][2].text, '1')
        self.assertEqual(root[2][3].tag, 'action')
        self.assertEqual(root[2][3].text, 'BUY')
        self.assertEqual(root[2][4].tag, 'username')
        self.assertEqual(root[2][4].text, 'cailan')
        self.assertEqual(root[2][5].tag, 'funds')
        self.assertEqual(root[2][5].text, '1000000')

        self.assertEqual(root[3].tag, 'systemEvent')
        self.assertEqual(root[3][0].tag, 'timestamp')
        self.assertEqual(root[3][1].tag, 'server')
        self.assertEqual(root[3][1].text, 'SERVER_1')
        self.assertEqual(root[3][2].tag, 'transactionNum')
        self.assertEqual(root[3][2].text, '1')
        self.assertEqual(root[3][3].tag, 'command')
        self.assertEqual(root[3][3].text, 'BUY')
        self.assertEqual(root[3][4].tag, 'username')
        self.assertEqual(root[3][4].text, 'cailan')
        self.assertEqual(root[3][5].tag, 'stockSymbol')
        self.assertEqual(root[3][5].text, 'HGU')

        self.assertEqual(root[4].tag, 'errorEvent')
        self.assertEqual(root[4][0].tag, 'timestamp')
        self.assertEqual(root[4][1].tag, 'server')
        self.assertEqual(root[4][1].text, 'SERVER_1')
        self.assertEqual(root[4][2].tag, 'transactionNum')
        self.assertEqual(root[4][2].text, '1')
        self.assertEqual(root[4][3].tag, 'command')
        self.assertEqual(root[4][3].text, 'BUY')
        self.assertEqual(root[4][4].tag, 'username')
        self.assertEqual(root[4][4].text, 'cailan')
        self.assertEqual(root[4][5].tag, 'stockSymbol')
        self.assertEqual(root[4][5].text, 'HGU')
        self.assertEqual(root[4][6].tag, 'errorMessage')
        self.assertEqual(root[4][6].text, 'Oops')

        self.assertEqual(root[5].tag, 'debugEvent')
        self.assertEqual(root[5][0].tag, 'timestamp')
        self.assertEqual(root[5][1].tag, 'server')
        self.assertEqual(root[5][1].text, 'SERVER_1')
        self.assertEqual(root[5][2].tag, 'transactionNum')
        self.assertEqual(root[5][2].text, '1')
        self.assertEqual(root[5][3].tag, 'command')
        self.assertEqual(root[5][3].text, 'BUY')
        self.assertEqual(root[5][4].tag, 'username')
        self.assertEqual(root[5][4].text, 'cailan')
        self.assertEqual(root[5][5].tag, 'stockSymbol')
        self.assertEqual(root[5][5].text, 'HGU')
        self.assertEqual(root[5][6].tag, 'debugMessage')
        self.assertEqual(root[5][6].text, 'Complete')
