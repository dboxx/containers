FROM alpine:3.14.2

RUN apk update \
 && apk upgrade \
 && apk add fio

WORKDIR /home
ENTRYPOINT ["fio"]
