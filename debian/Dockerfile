FROM debian:10.9-slim

ENV LANG=C.UTF-8

ENV SHELL=/bin/bash
ENV PAGER=less
ENV EDITOR=vim

ENV HOME=/root
ENV TERM=xterm-color
ENV PATH=/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin

ENV DEBIAN_VERSION="10"
ENV DEBIAN_RELEASE="${DEBIAN_VERSION}.7"
ENV DEBIAN_VERSION_NAME="buster"

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
      ifenslave \
      net-tools \
      coreutils \
      ssh-client \
      iputils-ping \
      bash-completion \
      ca-certificates \
      apt-transport-https \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""

COPY profile/bashrc       $HOME/.bashrc
COPY profile/bash_aliases $HOME/.bash_aliases
COPY profile/screenrc     $HOME/.screenrc

WORKDIR $HOME
