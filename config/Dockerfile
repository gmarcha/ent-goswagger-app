# syntax=docker/dockerfile:1

FROM golang:alpine

ARG PORT
ARG ENV

ENV ENV=${ENV}

RUN apk add make

WORKDIR /go/src/github.com/gmarcha/ent-goswagger-app

CMD make run ENV=${ENV}

EXPOSE ${PORT}
