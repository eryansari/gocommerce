FROM golang:1.15-alpine AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN go build -o .build/microservice-order ./microservice-order/cmd/main.go

FROM alpine:latest
ARG server=default
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta
ENV WORKDIR=/microservice-order
WORKDIR ${WORKDIR}
RUN mkdir ${WORKDIR}/log
COPY --from=builder /go/src/app/.build/microservice-order ${WORKDIR}/
COPY --from=builder /go/src/app/microservice-order/conf/config.yml ${WORKDIR}/conf/

CMD ./microservice-order ${server}