# syntax=docker/dockerfile:1

FROM golang:alpine

RUN apk add make

WORKDIR /app
COPY . .

CMD [ "make", "run" ]

EXPOSE 5000
