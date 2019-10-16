FROM dboxx/debian:v10.1

ENV HOME=/root
ENV TERM=xterm-color

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update -y \
 && apt-get install -yy -q \
      htop \
      iftop \
      xterm \
      screen \
      iperf3 \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND=""

COPY .bashrc       $HOME/.bashrc
COPY .bash_aliases $HOME/.bash_aliases
COPY .screenrc     $HOME/.screenrc

WORKDIR $HOME
