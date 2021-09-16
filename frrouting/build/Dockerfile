FROM frrouting/frr:v8.0.1

RUN sed -i "s/^bgpd=no/bgpd=yes/" /etc/frr/daemons

# expose bgp port
EXPOSE 179
