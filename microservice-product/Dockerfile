FROM golang:1.15-alpine AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN go build -o .build/microservice-product ./microservice-product/cmd/main.go

FROM alpine:latest
ARG server=default
RUN apk add --no-cache tzdata
ENV WORKDIR=/microservice-product
ENV TZ=Asia/Jakarta
WORKDIR ${WORKDIR}
RUN mkdir ${WORKDIR}/logs
COPY --from=builder /go/src/app/.build/microservice-product ${WORKDIR}/
COPY --from=builder /go/src/app/microservice-product/conf/config.yml ${WORKDIR}/conf/

CMD ./microservice-product ${server}