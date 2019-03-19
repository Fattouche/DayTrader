#!/bin/bash

bash /startup/wait-for-it.sh ${LOGGING_DB_IP}:${LOGGING_DB_PORT} --timeout=300

exec "$@"
