FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/Order-Service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/Order-Service /go/src/Order-Service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/Order-Service /usr/local/bin/Order-Service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["Order-Service"]
