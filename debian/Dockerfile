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
ENV APP_NAME=http-hello-test
ENV APP_DIR=/usr/local/bin

WORKDIR /tmp
COPY http-hello-test/main.go main.go

RUN go build -o ${APP_DIR}/${APP_NAME} \
       -ldflags="-linkmode external -extldflags '-static'" \
       ./main.go

# STAGE 2. Copy go application from stage 1 and execute it.
FROM debian:11.4-slim as application

ENV LANG=C.UTF-8

ENV SHELL=/bin/bash
ENV PAGER=less
ENV EDITOR=vim

ENV HOME=/root
ENV TERM=xterm-color
ENV PATH=/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin

ENV DEBIAN_VERSION="11"
ENV DEBIAN_RELEASE="${DEBIAN_VERSION}.4"
ENV DEBIAN_VERSION_NAME="bullseye"

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy --no-install-recommends \
      jq \
      tar \
      vim \
      git \
      bash \
      less \
      curl \
      vlan \
      htop \
      iftop \
      patch \
      screen \
      iperf3 \
      procps \
      telnet \
      gnupg2 \
      ethtool \
      locales \
      dnsutils \
      iproute2 \
      iptables \
      ifenslave \
      net-tools \
      coreutils \
      wireguard \
      ssh-client \
      iputils-ping \
      wireguard-tools \
      bash-completion \
      ca-certificates \
      apt-transport-https \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""

# app
ENV APP_NAME=http-hello-test
ENV APP_DIR=/usr/local/bin
ENV APP_BIN=${APP_DIR}/${APP_NAME}
COPY --from=builder ${APP_BIN} ${APP_BIN}

COPY profile/bashrc       $HOME/.bashrc
COPY profile/bash_aliases $HOME/.bash_aliases
COPY profile/screenrc     $HOME/.screenrc

WORKDIR $HOME
