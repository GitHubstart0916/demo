ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS builder
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -u github.com/swaggo/swag/cmd/swag

RUN mkdir -p /myapp
WORKDIR /myapp

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN swag init
RUN go build -o /app .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*
RUN cp "/usr/share/zoneinfo/Asia/Shanghai" "/etc/localtime" && echo "Asia/Shanghai" > "/etc/timezone"

RUN mkdir -p /data /logs
WORKDIR /data
EXPOSE 8000

CMD /app -listen 0.0.0.0:8000
