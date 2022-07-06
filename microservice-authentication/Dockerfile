FROM golang:1.15-alpine AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ .
RUN go build -o .build/microservice-authentication ./microservice-authentication/cmd/main.go

FROM alpine:latest
ARG server=default
RUN apk add --no-cache tzdata
ENV WORKDIR=/microservice-authentication
ENV TZ=Asia/Jakarta
WORKDIR ${WORKDIR}
RUN mkdir ${WORKDIR}/logs
COPY --from=builder /go/src/app/.build/microservice-authentication ${WORKDIR}/
COPY --from=builder /go/src/app/microservice-authentication/conf/config.yml ${WORKDIR}/conf/

CMD ./microservice-authentication ${server}