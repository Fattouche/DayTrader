FROM golang

WORKDIR /startup
COPY . /startup

RUN go get -insecure github.com/rainycape/memcache
RUN go get -insecure github.com/go-sql-driver/mysql
RUN go get -insecure google.golang.org/grpc
RUN go get -insecure google.golang.org/grpc/reflection
RUN go get -insecure -u github.com/Fattouche/DayTrader/golang/protobuff
RUN go get -insecure github.com/pkg/profile
RUN go get -insecure github.com/mailru/easyjson/...

WORKDIR /go/src/

ENTRYPOINT ["bash", "/startup/entrypoint.sh"]