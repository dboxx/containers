FROM debian:buster-slim

# OS
ENV DEBIAN_VERSION  "buster"
ENV DEBIAN_FRONTEND "noninteractive"

# GlusterFS configuration
ENV GFS_VERSION "6"
ENV GFS_RELEASE "6.1"

# GlusterFS APT configuration
ENV GFS_APT_URL "https://download.gluster.org/pub/gluster/glusterfs"
ENV GFS_APT_KEY "${GFS_APT_URL}/${GFS_VERSION}/rsa.pub"
ENV GFS_APT_SRC "${GFS_APT_URL}/${GFS_VERSION}/${GFS_RELEASE}/Debian/${DEBIAN_VERSION}/amd64/apt"

RUN apt-get update -y \
 && apt-get install -yy -q curl gnupg2 \
 && curl -sSk "${GFS_APT_KEY}" | apt-key add - \
 && echo "deb [arch=amd64] ${GFS_APT_SRC} ${DEBIAN_VERSION} main" > /etc/apt/sources.list.d/gluster.list \
 && cat /etc/apt/sources.list.d/gluster.list \
 && apt-get update -y \
 && apt-get install -yy -q glusterfs-server \
 && apt-get remove -y --purge curl gnupg2 \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND ""
