FROM debian:10.1-slim

ENV SHELL=/bin/bash
ENV EDITOR=vim
ENV PAGER=less
ENV PATH=/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin

ENV DEBIAN_VERSION="10"
ENV DEBIAN_VERSION_NAME="buster"

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy -q \
      jq \
      vim \
      bash \
      less \
      curl \
      vlan \
      telnet \
      gnupg2 \
      ethtool \
      locales \
      dnsutils \
      iproute2 \
      ifenslave \
      net-tools \
      coreutils \
      iputils-ping \
      bash-completion \
      ca-certificates \
      command-not-found \
      apt-transport-https \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""

WORKDIR /

ENTRYPOINT ["/bin/bash"]
