FROM dboxx/qemu:v3.1

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get install -yy -q \
      libvirt-dbus \
      libvirt-daemon \
      libvirt-clients \
      libvirt-daemon-system \
      libvirt-daemon-driver-storage-zfs \
      libvirt-daemon-driver-storage-rbd \
      libvirt-daemon-driver-storage-gluster \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""
