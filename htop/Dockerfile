FROM alpine:3.14

RUN apk update \
 && apk upgrade \
 && apk add htop

COPY htoprc /etc/htoprc

ENV PATH "/usr/bin"

WORKDIR /

ENTRYPOINT ["htop"]
