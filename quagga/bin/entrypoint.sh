#!/bin/sh

if [ $1 = "quagga" ]; then
    echo "Starting loop ... "
    while [ 1 ]; do
        sleep 60
    done
elif [ $1 = "vtysh" ]; then
    echo "Starting vtysh ... "
    $@
else
  echo "Starting $1 daemon ... "
  /usr/sbin/$1 -f /etc/quagga/$1.conf
fi
