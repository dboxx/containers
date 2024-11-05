FROM postgres:17.0-bookworm

ENV PATH "/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/sbin"

RUN apt update \
 && apt install -y iputils-ping vim curl netcat-openbsd iproute2 postgresql-contrib

WORKDIR /
