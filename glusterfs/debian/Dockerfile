FROM debian:9.6

ENV DEBIAN_FRONTEND=noninteractive
ENV LD_LIBRARY_PATH=/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy -q \
              git \
              uuid-dev \
              make \
              automake \
              autoconf \
              libtool \
              flex \
              bison \
              pkg-config \
              libssl-dev \
              libxml2-dev \
              python-dev \
              libaio-dev \
              libibverbs-dev \
              librdmacm-dev \
              libreadline-dev \
              liblvm2-dev \
              libglib2.0-dev \
              liburcu-dev \
              libcmocka-dev \
              libsqlite3-dev \
              libacl1-dev \
#           jq \
#           gcc \
#           tar \
#           uuid \
#           curl \
#           flex \
#           make \
#           cmake \
#           bison \
#           python \
#           libtool \
#           autoconf \
#           binutils \
#           automake \
#           rpcbind \
#           pkg-config \
#           acl-dev \
#           libc-dev \
#           musl-dev \
#           libnfs-dev \
#           libssh-dev \
#           libfuse-dev \
#           liburcu-dev \
#           libcrypto++-dev \
#           python-dev \
#           libxml2-dev \
#           libc-dev-bin \
#           libtirpc-dev \
#           libsqlite3-dev \
#           libreadline-dev \
#           libcrossguid-dev \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && rm -rf /var/cache/apt

ENV APP_DIR=/src/github.com/gluster/glusterfs

#ENV GLUSTERFS_BRANCH="master"
ENV GLUSTERFS_BRANCH="release-5"

RUN git clone -b "${GLUSTERFS_BRANCH}" --single-branch --depth 1 https://github.com/gluster/glusterfs.git ${APP_DIR}

RUN cd ${APP_DIR} \
 && ./autogen.sh \
 && ./configure --enable-static \
                --disable-glupy \
                --disable-syslog \
                --disable-events \
                --disable-ibverbs \
                --disable-selinux \
                --disable-tiering \
                --disable-ec-dynamic \
                --without-libtirpc \
 && make \
 && make install

#CMD /usr/local/sbin/glusterd -p /run/glusterd.pid --debug --log-level DEBUG
