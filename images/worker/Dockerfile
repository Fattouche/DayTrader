FROM python

RUN mkdir /service
WORKDIR /service

COPY ./worker/requirements.txt .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

COPY ./worker .

ENTRYPOINT ["bash", "./entrypoint.sh"]

