#!/bin/bash

cd day_trader
#If you change the schema you need to run this
#easyjson -all main.go  

bash /startup/wait-for-it.sh ${DAYTRADER_DB_IP}:3306 --timeout=300
bash /startup/wait-for-it.sh ${DAYTRADER_CACHE_IP}:${DAYTRADER_CACHE_PORT} --timeout=300
bash /startup/wait-for-it.sh ${LOGGING_LB_IP}:${LOGGING_LB_PORT} --timeout=300



exec "$@"
