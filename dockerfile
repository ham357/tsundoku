FROM golang:latest

COPY entrypoint.sh /var/tmp
CMD bash -E /var/tmp/entrypoint.sh && /bin/bash

ENV SRC_DIR=/go/1.15.2/src/github.com/ham357/tsundoku/api

WORKDIR $SRC_DIR

ADD ./api $SRC_DIR

RUN cd /go/1.15.2/src/;

RUN go get -u github.com/go-sql-driver/mysql \
  && go get -u github.com/labstack/echo/... \
  && go get -u github.com/gorilla/mux \
  && go get -u github.com/jinzhu/gorm \
  && go get -u github.com/gin-contrib/cors \
  && go get -u gopkg.in/ini.v1 \
  && go get -u github.com/cosmtrek/air \
  && go get -u github.com/go-delve/delve/cmd/dlv \
  && go get -u firebase.google.com/go/ \
  && go get -u github.com/valyala/fasthttp

CMD air -c .air.toml
