# fio - flexible I/O tester

## Usage
```
docker run --net=host --pid=host --privileged -v "`pwd`:/home:rw" --rm -it dboxx/fio
```

## Tests

### Write 64k
```
docker run --net=host --pid=host --privileged -v "`pwd`:/home:rw" --rm -it dboxx/fio -name=write-64k -filename=write-64k.fio -rw=write -group_reporting -direct=1 -iodepth 32 -thread -ioengine=psync -numjobs=16 -bs=64k -size=1G
```

### Read 64k
```
docker run --net=host --pid=host --privileged -v "`pwd`:/home:rw" --rm -it dboxx/fio -name=read-64k -filename=read-64k.fio -rw=read -group_reporting -direct=1 -iodepth 32 -thread -ioengine=psync -numjobs=16 -bs=64k -size=1G
```
