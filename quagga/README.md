# quagga
Quagga is a network routing software

```bash
> make local-up
> make local-sh
> show ip route
```

## Run zebra
```
docker run \
  --name=zebra \
  --net=host \
  --pid=host \
  --privileged \
  -v "`pwd`/fixtures/matchbox-bgp/etc/quagga:/etc/quagga:rw" \
  -v "`pwd`/fixtures/matchbox-bgp/var/run/quagga:/var/run/quagga:rw" \
  --rm \
  -d \
  dboxx/quagga zebra
```

## Run bgpd
```
docker run \
  --name=bgpd \
  --net=host \
  --pid=host \
  --privileged \
  -v "`pwd`/fixtures/matchbox-bgp/etc/quagga:/etc/quagga:rw" \
  -v "`pwd`/fixtures/matchbox-bgp/var/run/quagga:/var/run/quagga:rw" \
  --rm \
  -d \
  dboxx/quagga bgpd
```
