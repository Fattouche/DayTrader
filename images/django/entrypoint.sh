#!/bin/bash

bash /service/images/django/wait-for-it.sh db:3306 --timeout=300

bash /service/images/django/wait-for-it.sh cache:11211 --timeout=300

exec "$@"