FROM debian:buster-slim

ENV VERSION "5"
ENV RELEASE "5.5"
ENV DEBIAN_VERSION "buster"
ENV GLUSTER_APT_KEY "https://download.gluster.org/pub/gluster/glusterfs/$VERSION/rsa.pub"
ENV GLUSTER_APT_SOURCE "https://download.gluster.org/pub/gluster/glusterfs/$VERSION/$RELEASE/Debian/$DEBIAN_VERSION/amd64/apt"

RUN export DEBIAN_FRONTEND=noninteractive \
 && apt-get update -y \
 && apt-get install -yy -q curl gnupg2 \
 && curl -sSk $GLUSTER_APT_KEY | apt-key add - \
 && echo deb [arch=amd64] $GLUSTER_APT_SOURCE $DEBIAN_VERSION main > /etc/apt/sources.list.d/gluster.list \
 && cat /etc/apt/sources.list.d/gluster.list \
 && apt-get update -y \
 && apt-get remove -y --purge curl gnupg2 \
 && apt-get install -yy -q glusterfs-server \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*
