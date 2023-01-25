FROM golang AS builder

WORKDIR /go/src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/ordersystem


FROM alpine:3.17

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz


COPY --from=builder /go/src/cmd/ordersystem/.env .
COPY --from=builder /go/src/main .
