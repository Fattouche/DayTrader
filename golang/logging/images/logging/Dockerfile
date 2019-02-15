FROM golang

#ENV http_proxy 'http://192.168.1.1:3128'
#ENV https_proxy 'https://192.168.1.1:3128'

WORKDIR /startup
COPY . /startup

RUN go get -insecure github.com/go-sql-driver/mysql
RUN go get -insecure google.golang.org/grpc
RUN go get -insecure google.golang.org/grpc/reflection
RUN go get -insecure -u github.com/Fattouche/DayTrader/golang/protobuff

WORKDIR /go/src/logging

ENTRYPOINT ["bash", "/startup/entrypoint.sh"]