FROM golang

WORKDIR /startup
COPY . /startup

RUN go get github.com/bradfitz/gomemcache/memcache
RUN go get github.com/go-sql-driver/mysql
RUN go get google.golang.org/grpc
RUN go get google.golang.org/grpc/reflection
RUN go get -u github.com/Fattouche/DayTrader/golang/protobuff
RUN go get github.com/pkg/profile
RUN go get github.com/mailru/easyjson/...

WORKDIR /go/src/

ENTRYPOINT ["bash", "/startup/entrypoint.sh"]