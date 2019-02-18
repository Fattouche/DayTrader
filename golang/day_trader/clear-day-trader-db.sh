#!/bin/bash
if [ "$#" -ne 1 ]; then
    echo "Usage: clear-day-trader-db.sh [mysql-port-number]"
    exit 1
fi
mysql -u root -P $1 <<QUERY_INPUT
use daytrader;
DELETE FROM Buy;
DELETE FROM Buy_Trigger;
DELETE FROM Sell;
DELETE FROM Sell_Trigger;
DELETE FROM User;
DELETE FROM User_Stock;
QUERY_INPUT

echo "Done"