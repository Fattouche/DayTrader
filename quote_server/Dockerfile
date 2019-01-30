FROM golang
RUN mkdir /service
COPY . /service
WORKDIR /service
EXPOSE 4442

ENTRYPOINT [ "go run server.go" ]
