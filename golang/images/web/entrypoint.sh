#!/bin/bash

cd day_trader
easyjson -all cache.go

bash /startup/wait-for-it.sh db:3306 --timeout=300

bash /startup/wait-for-it.sh cache:11211 --timeout=300



exec "$@"
