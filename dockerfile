FROM golang:latest

ENV SRC_DIR=/go/1.15.2/src/github.com/ham357/tsundoku/api

WORKDIR $SRC_DIR

ADD ./api $SRC_DIR

RUN cd /go/1.15.2/src/;

RUN go get -u github.com/labstack/echo/...

ENTRYPOINT ["go", "run", "server.go"]

