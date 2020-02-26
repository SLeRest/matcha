FROM  golang:1.13.4-alpine

WORKDIR /usr/src/api

RUN apk update && apk add --no-cache git

RUN go get -u github.com/gorilla/mux
RUN go get gopkg.in/go-playground/validator.v9

COPY . /usr/src/api
