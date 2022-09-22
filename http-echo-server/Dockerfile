# STAGE 1. Build go application.
FROM golang:1.19.1-alpine3.16 as builder

RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen
RUN apk add --update --no-cache ca-certificates gcc make binutils libc-dev musl-dev

# golang
ENV GOPATH=/go
ENV GOROOT=/usr/local/go

# app
ENV APP_DIR=/app
ENV APP_NAME=http-echo-server

WORKDIR ${APP_DIR}
COPY main.go main.go

# linux/amd64 (x86_64)
RUN GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}-amd64 -ldflags="-linkmode external -extldflags '-static'" ./main.go

# linux ppc64le
RUN GOOS=linux GOARCH=ppc64le go build -o ${APP_NAME}-ppc64le ./main.go

# linux/arm64 (aarch64)
RUN GOOS=linux GOARCH=arm64 go build -o ${APP_NAME}-arm64 ./main.go

# linux/386 (i386)
RUN GOOS=linux GOARCH=386 go build -o ${APP_NAME}-386 ./main.go

# STAGE 2. Copy go application from stage 1 and execute it.
FROM alpine:3.16 as application

RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen
RUN apk add --update --no-cache ca-certificates curl jq

# app
ENV APP_NAME=http-echo-server

COPY --from=builder /app/${APP_NAME}-* /usr/local/bin/

WORKDIR /usr/local/bin
COPY entrypoint.sh ${APP_NAME}
RUN ln -s ${APP_NAME}-amd64 ${APP_NAME}-x86_64
RUN ln -s ${APP_NAME}-arm64 ${APP_NAME}-aarch64
RUN ln -s ${APP_NAME}-386 ${APP_NAME}-i386

WORKDIR /
CMD ${APP_NAME}
