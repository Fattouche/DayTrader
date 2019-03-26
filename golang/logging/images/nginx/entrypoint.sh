#!/bin/bash

cd /service

# bash ./wait-for-it.sh ${LOGGING_SERVER_IP}:${LOGGING_SERVER_PORT} --timeout=300
bash ./wait-for-it.sh 192.168.1.240:40000 --timeout=300
bash ./wait-for-it.sh 192.168.1.187:40000 --timeout=300

#Uncomment to scale up
#bash ./wait-for-it.sh logging_1:40000 --timeout=300
#bash ./wait-for-it.sh logging_2:40000 --timeout=300
#bash ./wait-for-it.sh logging_3:40000 --timeout=300

exec nginx -g 'daemon off;'
