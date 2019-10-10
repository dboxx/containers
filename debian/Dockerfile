FROM debian:10.1-slim

ENV DEBIAN_VERSION="10"
ENV DEBIAN_VERSION_NAME="buster"

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy -q \
      jq \
      vim \
      curl \
      telnet \
      gnupg2 \
      dnsutils \
      iproute2 \
      net-tools \
      iputils-ping \
      ca-certificates \
      apt-transport-https \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""
