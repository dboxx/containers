FROM dboxx/debian

# GlusterFS configuration
ENV GFS_VERSION="6"
ENV GFS_RELEASE="6.5"

# GlusterFS APT configuration
ENV GFS_APT_URL="https://download.gluster.org/pub/gluster/glusterfs"
ENV GFS_APT_KEY="${GFS_APT_URL}/${GFS_VERSION}/rsa.pub"
ENV GFS_APT_SRC="${GFS_APT_URL}/${GFS_VERSION}/${GFS_RELEASE}/Debian/${DEBIAN_VERSION_NAME}/amd64/apt"

ENV DEBIAN_FRONTEND="noninteractive"

RUN curl -sSk "${GFS_APT_KEY}" | apt-key add - \
 && echo "deb [arch=amd64] ${GFS_APT_SRC} ${DEBIAN_VERSION_NAME} main" > /etc/apt/sources.list.d/gluster.list \
 && cat /etc/apt/sources.list.d/gluster.list \
 && apt-get update -y \
 && apt-get install -yy -q glusterfs-server \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""
