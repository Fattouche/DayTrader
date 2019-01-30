FROM golang
RUN mkdir /service
COPY . /service
WORKDIR /service
EXPOSE 4442
RUN chmod +x server.go
ENTRYPOINT [ "go",  "run" , "server.go" ]
