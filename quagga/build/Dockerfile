FROM alpine:3.14.2

RUN apk update \
 && apk upgrade \
 && apk add quagga

COPY bin/entrypoint.sh /usr/local/bin/entrypoint.sh
COPY etc/quagga/zebra.conf /etc/quagga/zebra.conf
COPY etc/quagga/bgpd.conf /etc/quagga/bgpd.conf

# expose bgp port
EXPOSE 179

# expose quagga mgmt ports
EXPOSE 2601 2602 2603 2604 2605 2606

WORKDIR /home
ENTRYPOINT ["entrypoint.sh"]
