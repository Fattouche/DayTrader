#!/bin/bash

bash ./wait-for-it.sh queue:6379 --timeout=300

exec "$@"
