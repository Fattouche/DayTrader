FROM golang

WORKDIR /startup
COPY . /startup

RUN go get github.com/go-sql-driver/mysql
RUN go get google.golang.org/grpc
RUN go get google.golang.org/grpc/reflection
RUN go get -u github.com/Fattouche/DayTrader/golang/protobuff

WORKDIR /service

ENTRYPOINT ["/bin/bash"]
# ENTRYPOINT ["bash", "/startup/entrypoint.sh"]