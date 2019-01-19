FROM python:3.6.5

RUN mkdir /service
WORKDIR /service

COPY ./day_trader/requirements.txt .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

ENTRYPOINT ["bash", "./images/django/entrypoint.sh"]