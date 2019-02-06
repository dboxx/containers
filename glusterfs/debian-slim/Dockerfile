FROM debian:9.6 as builder

ENV DEBIAN_FRONTEND=noninteractive
ENV LD_LIBRARY_PATH=/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy -q \
            jq \
            gcc \
            git \
            tar \
            uuid \
            curl \
            flex \
            make \
            cmake \
            bison \
            python \
            libtool \
            autoconf \
            binutils \
            automake \
            rpcbind \
            pkg-config \
            acl-dev \
            libc-dev \
            musl-dev \
            uuid-dev \
            libnfs-dev \
            libssh-dev \
            libfuse-dev \
            liburcu-dev \
            libcrypto++-dev \
            python-dev \
            libxml2-dev \
            libc-dev-bin \
            libtirpc-dev \
            libsqlite3-dev \
            libreadline-dev \
            libcrossguid-dev \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && rm -rf /var/cache/apt

# checking for SQLITE... no
# configure: error: pass --disable-tiering to build without sqlite

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

FROM debian:9.6-slim as application

ENV DEBIAN_FRONTEND=noninteractive
ENV LD_LIBRARY_PATH=/usr/local/lib:/usr/lib/x86_64-linux-gnu:/lib/x86_64-linux-gnu

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy -q \
            python \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && rm -rf /var/cache/apt

COPY --from=builder  /sbin/mount.glusterfs  /sbin/mount.glusterfs

COPY --from=builder  /usr/local/lib   /usr/local/lib
COPY --from=builder  /usr/local/bin   /usr/local/bin
COPY --from=builder  /usr/local/sbin  /usr/local/sbin

COPY --from=builder  /lib/x86_64-linux-gnu      /lib/x86_64-linux-gnu
COPY --from=builder  /usr/lib/x86_64-linux-gnu  /usr/lib/x86_64-linux-gnu

COPY --from=builder  /usr/local/etc/glusterfs     /usr/local/etc/glusterfs
COPY --from=builder  /usr/local/include/glusterfs /usr/local/include/glusterfs
COPY --from=builder  /usr/local/libexec/glusterfs /usr/local/libexec/glusterfs
COPY --from=builder  /var/lib/glusterd            /var/lib/glusterd
COPY --from=builder  /var/log/glusterfs           /var/log/glusterfs

CMD /usr/local/sbin/glusterd -p /run/glusterd.pid --debug --log-level DEBUG
