#!/bin/bash
if [ "$#" -ne 1 ]; then
    echo "Usage: clear-logging-db.sh [mysql-port-number]"
    exit 1
fi
mysql -u root -P $1 <<QUERY_INPUT
use logging;
TRUNCATE TABLE UserCommandLog;
TRUNCATE TABLE AccountTransactionLog;
TRUNCATE TABLE QuoteServerLog;
TRUNCATE TABLE SystemEventLog;
TRUNCATE TABLE DebugEventLog;
TRUNCATE TABLE ErrorEventLog;
QUERY_INPUT

echo "Done"
