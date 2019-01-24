FROM python

RUN mkdir /startup
COPY ./requirements.txt /startup
WORKDIR /startup

RUN pip install --upgrade pip
RUN pip install -r /startup/requirements.txt

COPY . /startup
WORKDIR /day_trader

ENTRYPOINT ["bash", "/startup/entrypoint.sh"]
