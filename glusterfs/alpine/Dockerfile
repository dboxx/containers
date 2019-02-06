FROM alpine:3.8

RUN apk update \
 && apk upgrade \
 && apk add jq\
            gcc \
            git \
            tar \
            curl \
            file \
            make \
            flex \
            bison \
            pkgconf \
            libc-dev \
            musl-dev \
            binutils \
            automake \
            autoconf \
            libtool \
            rpcgen \
            openssl-dev \
            util-linux-dev \
            acl-dev \
            python2 \
            python2-dev \
            libxml2-dev \
            argp-standalone \
            libtirpc-dev \
            userspace-rcu-dev \
            readline-dev \
            libnfs-dev \
            nfs-utils-dev \
            sqlite-dev

# checking for SQLITE... no
# configure: error: pass --disable-tiering to build without sqlite

ENV APP_DIR=/src/github.com/gluster/glusterfs

ENV GLUSTERFS_BRANCH="v5.3"
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
# && make
 && true

