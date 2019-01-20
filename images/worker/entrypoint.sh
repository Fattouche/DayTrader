#!/bin/bash

chmod +x ./wait-for-it.sh

bash ./wait-for-it.sh queue:6379 --timeout=300

exec "$@"