from django.db import models


class BaseLog(models.Model):
    log_type = models.CharField(max_length=32)
    timestamp = models.DateTimeField(auto_now_add=True)
    server = models.CharField(max_length=8)
    transaction_num = models.BigIntegerField()


class UserCommandLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.DecimalField(decimal_places=2, max_digits=32, null=True)


class QuoteServerLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    price = models.DecimalField(decimal_places=2, max_digits=32)
    stock_symbol = models.CharField(max_length=3)
    username = models.CharField(max_length=128)
    quote_server_time = models.BigIntegerField()
    # TODO: what is crypto key length?
    crypto_key = models.CharField(max_length=256)


class AccountTransactionLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    action = models.CharField(max_length=32)
    username = models.CharField(max_length=64)
    funds = models.DecimalField(decimal_places=2, max_digits=32)


class SystemEventLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.DecimalField(decimal_places=2, max_digits=32, null=True)


class ErrorEventLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.DecimalField(decimal_places=2, max_digits=32, null=True)
    error_message = models.CharField(max_length=512)


class DebugEventLog(models.Model):
    base_log = models.OneToOneField(BaseLog, on_delete=models.CASCADE)
    command = models.CharField(max_length=16)
    username = models.CharField(max_length=64)
    stock_symbol = models.CharField(max_length=3)
    filename = models.FilePathField(path="/dumplog_output")
    funds = models.DecimalField(decimal_places=2, max_digits=32, null=True)
    debug_message = models.CharField(max_length=512)
