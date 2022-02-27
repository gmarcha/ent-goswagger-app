# syntax=docker/dockerfile:1

FROM golang:alpine

ARG PORT

RUN apk add make

WORKDIR /go/src/github.com/gmarcha/ent-goswagger-app

CMD [ "make", "run" ]

EXPOSE ${PORT}
