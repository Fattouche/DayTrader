# Generated by Django 2.1.5 on 2019-01-27 23:01

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('exchange', '0002_auto_20190127_2051'),
    ]

    operations = [
        migrations.AddField(
            model_name='selltrigger',
            name='stock_reserved_amount',
            field=models.PositiveIntegerField(default=0),
        ),
    ]