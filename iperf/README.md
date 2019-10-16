# iftop
Display bandwidth usage on an interface by host

## Run server
```
docker run --net=host --pid=host --privileged --rm -it dboxx/iperf -s
```

## Run client
```
docker run --net=host --pid=host --privileged --rm -it dboxx/iperf -c [iperf.server.ip.address]
```
