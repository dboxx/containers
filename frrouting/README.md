# frrouting

## FRRouting Project
FRR is free software that implements and manages various IPv4 and IPv6 routing protocols.
It runs on nearly all distributions of Linux and BSD and supports all modern CPU architectures.

- https://frrouting.org/
- https://github.com/kubeovn/kube-ovn/blob/master/docs/bgp.md
- https://ltomasbo.wordpress.com/2021/02/04/ovn-bgp-agent-testing-setup/
- http://www.openvswitch.org/support/ovscon2020/slides/OVS-CONF-2020-OVN-WITH-DYNAMIC-ROUTING.pdf

## MetalLB config
```bash
apiVersion: v1
kind: ConfigMap
data:
  config: |
    peers:
    - my-asn: 64501
      peer-address: 10.0.0.1
      peer-asn: 64500
    address-pools:
    - addresses:
      - 10.0.0.0/24
      name: default
      protocol: bgp
```

## FRRouting bgp config

### Route-Map config
```
route-map ACCEPT_METALLB_IP permit 10
 match ip address prefix-list INTERNAL_METALLB_IP_Z1
!
route-map ACCEPT_METALLB_IP permit 20
 match ip address prefix-list EXTERNAL_METALLB_IP_Z1
!
route-map ADVERTISE_ONLY_DEFAULT permit 10
 match ip address prefix-list DEFAULT_ONLY
!
route-map LOOPBACK permit 10
 match interface lo
```

### Router config
```
router bgp 64500
 bgp router-id 10.0.0.1
 no bgp default ipv4-unicast
 neighbor METALLB peer-group
 neighbor METALLB remote-as external
 neighbor METALLB disable-connected-check
 bgp listen range 10.0.0.0/24 peer-group METALLB
 !
 address-family ipv4 unicast
  redistribute connected route-map LOOPBACK
  neighbor METALLB activate
  neighbor METALLB route-map ACCEPT_METALLB_IP in
  neighbor METALLB route-map ADVERTISE_ONLY_DEFAULT out
 exit-address-family
!
ip prefix-list DEFAULT_ONLY seq 10 permit 0.0.0.0/0
ip prefix-list INTERNAL_METALLB_IP_Z1 seq 10 permit 10.0.0.0/24 ge 32
ip prefix-list EXTERNAL_METALLB_IP_Z1 seq 10 permit 10.0.0.0/24 ge 32
```
