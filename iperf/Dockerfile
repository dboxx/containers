FROM alpine:3.13

RUN apk update \
 && apk upgrade \
 && apk add iperf3

ENV PATH "/usr/bin"

WORKDIR /

ENTRYPOINT ["iperf3"]
