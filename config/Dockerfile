# syntax=docker/dockerfile:1

FROM golang:alpine

ARG PORT

RUN apk add make

WORKDIR /app
COPY . .

CMD [ "make", "run" ]

EXPOSE ${PORT}
