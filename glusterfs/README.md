# GlusterFS

## Volume creation

Run glusterd
```
 > glusterd
```

Create some directory for gluster volume
```
 > mkdir /tmp/test-vol
```

Create and start gluster volume
```
 > gluster volume create test-vol 172.17.0.2:/tmp/test-vol
 > gluster volume start test-vol
```

Check this volume in gluster
```
 > gluster volume list
 > gluster volume info test-vol
```

Create some directory for mount
```
mkdir /tmp/test-mnt
```

Mount test-vol to this directory
```
mount -t glusterfs 172.17.0.2:/tmp/test-vol /tmp/test-mnt
```

Test this volume
```
echo "test" > /tmp/test-mnt/test
cat /tmp/test-vol/test
```
