# STAGE 1. Build go application.
FROM golang:1.19.0-alpine3.16 as builder

RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen

RUN apk update \
 && apk upgrade \
 && apk add ca-certificates gcc make libc-dev musl-dev binutils git curl jq \
 && rm -rf /var/cache/apk/*

# golang
ENV GOPATH=/go
ENV GOROOT=/usr/local/go

# app
ENV APP_NAME=go-hello
ENV APP_DIR=/app

WORKDIR ${APP_DIR}
COPY main.go main.go

RUN go build -o ${APP_NAME} \
       -ldflags="-linkmode external -extldflags '-static'" \
       ./main.go

# STAGE 2. Copy go application from stage 1 and execute it.
FROM scratch as application

# app
ENV APP_NAME=go-hello
ENV APP_DIR=/app
ENV APP_BIN=${APP_DIR}/${APP_NAME}

COPY --from=builder ${APP_BIN} ${APP_BIN}

WORKDIR /
CMD ${APP_BIN}
