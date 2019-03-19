#!/bin/bash

cd /service

#Uncomment to scale up
 bash ./wait-for-it.sh 192.168.1.157:41000 --timeout=300
 bash ./wait-for-it.sh 192.168.1.194:41000 --timeout=300
 bash ./wait-for-it.sh 192.168.1.191:41000 --timeout=300

/usr/local/bin/envoy -c /etc/envoy/envoy.yaml
