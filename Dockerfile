FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/order-service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/order-service /go/src/order-service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/order-service /usr/local/bin/order-service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["order-service"]
