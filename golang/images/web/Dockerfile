FROM golang

WORKDIR /startup
COPY . /startup
RUN go get .

WORKDIR /day_trader

ENTRYPOINT ["bash", "/startup/entrypoint.sh"]