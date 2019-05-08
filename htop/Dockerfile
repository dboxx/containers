FROM alpine:3.9

RUN apk update \
 && apk upgrade \
 && apk add htop

COPY htoprc /etc/htoprc

ENV PATH "/bin:/usr/bin"

WORKDIR /

ENTRYPOINT ["htop"]
