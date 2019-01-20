#!/bin/bash

chmod +x /service/images/wait-for-it.sh

bash /service/images/wait-for-it.sh db:3306 --timeout=300

bash /service/images/wait-for-it.sh cache:11211 --timeout=300

python /service/day_trader/manage.py makemigrations
python /service/day_trader/manage.py migrate

exec "$@"