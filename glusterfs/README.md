## Create volume
```
glusterd
rm -rf /tmp/test-vol
mkdir -p /tmp/test-vol
gluster volume create test-vol 172.17.0.2:/tmp/test-vol
gluster volume list
mkdir /tmp/test-mnt
mount -t glusterfs 172.17.0.2:/tmp/test-vol /tmp/test-mnt
echo "test" > /tmp/test-mnt/test
cat /tmp/test-vol/test
```
