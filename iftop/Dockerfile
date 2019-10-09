FROM alpine:3.10

RUN apk update \
 && apk upgrade \
 && apk add iftop

ENV PATH "/usr/sbin"

WORKDIR /

ENTRYPOINT ["iftop"]
