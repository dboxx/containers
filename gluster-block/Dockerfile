FROM dboxx/glusterfs:v7.0

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get install -yy \
      gcc \
      gdb \
      file \
      kmod \
      make \
      cmake \
      zlib1g \
      libtool \
      autoconf \
      automake \
      libc-dev \
      uuid-dev \
      librados2 \
      libgfapi0 \
      python-dev \
      pkg-config \
      librbd-dev \
      zlib1g-dev \
      libkmod-dev\
      libnl-3-dev \
      libglib2.0-0 \
      targetcli-fb \
      libtirpc-dev \
      libpython-dev \
      libjson-c-dev \
      python-gobject \
      libglib2.0-dev \
      libglusterfs-dev \
      libnl-genl-3-dev \
      python-rtslib-fb \
      python-setuptools \
      libgoogle-perftools-dev\
 && apt-get remove -y --purge \
      targetcli-fb \
      python-rtslib-fb \
      python3-configshell-fb \
      python3-rtslib-fb \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && mkdir -p /etc/target /usr/lib/systemd/system \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

WORKDIR /usr/src
RUN git clone -b v0.4 https://github.com/gluster/gluster-block.git
RUN git clone https://github.com/open-iscsi/tcmu-runner.git
RUN git clone https://github.com/open-iscsi/targetcli-fb.git
RUN git clone https://github.com/open-iscsi/rtslib-fb.git
RUN git clone https://github.com/open-iscsi/configshell-fb.git

WORKDIR /usr/src/tcmu-runner
RUN cmake -DSUPPORT_SYSTEMD=ON -DCMAKE_INSTALL_PREFIX=/usr -Dwith-rbd=false -Dwith-qcow=false -Dwith-zbc=false -Dwith-fbo=false \
 && make install

WORKDIR /usr/src/targetcli-fb
RUN ./setup.py install

WORKDIR /usr/src/rtslib-fb
RUN ./setup.py install \
 && cp systemd/target.service /usr/lib/systemd/system/target.service

WORKDIR /usr/src/configshell-fb
RUN ./setup.py install

RUN ln -s /bin/ps /usr/bin/ps \
 && ln -s /usr/local/bin/targetctl  /usr/bin/targetctl \
 && ln -s /usr/local/bin/targetcli  /usr/bin/targetcli \
 && ln -s /usr/local/bin/targetclid /usr/bin/targetclid

WORKDIR /usr/src/gluster-block
RUN ./autogen.sh \
 && ./configure --prefix=/usr --with-systemddir=/usr --enable-tirpc=no \
 && make -j install \
 && cp systemd/* /usr/lib/systemd/system/

ENV DEBIAN_FRONTEND=""

RUN systemctl enable target
RUN systemctl enable tcmu-runner
RUN systemctl enable gluster-block-target
RUN systemctl enable gluster-blockd

WORKDIR $HOME
