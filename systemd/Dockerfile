FROM dboxx/debian:v10.9

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get upgrade -y \
 && apt-get install -yy \
      dbus \
      systemd \
      systemd-sysv \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""

CMD ["/sbin/init"]
