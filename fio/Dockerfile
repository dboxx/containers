FROM alpine:3.10

RUN apk update \
 && apk upgrade \
 && apk add fio

ENV PATH "/usr/bin"

WORKDIR /home

ENTRYPOINT ["fio"]
