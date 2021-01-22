FROM dboxx/debian:v10.7

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get install -yy -q \
      qemu \
      qemu-kvm \
      qemu-efi \
      qemu-utils \
      qemu-block-extra \
      qemu-guest-agent \
      qemu-system-data \
      qemu-system-common \
      qemu-system-x86 \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""
