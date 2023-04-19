## Fstab Configuration

Volume configuration

### Install Dependency
```
sudo apt-get -y update
sudo apt-get -y install libblkid1
```

### Backup & Configure `/etc/fstab`
* Copy / Backup `/etc/fstab`
```
sudo cp /etc/fstab /etc/fstab_old
```

* Show `/etc/fstab` file
```
LABEL=cloudimg-rootfs   /                     ext4     defaults    0 0
```

* Find disk UUID or view block partition
```
sudo lsblk -f
-------
NAME   FSTYPE LABEL           UUID                                 MOUNTPOINT
vda                                                                
└─vda1 ext4   cloudimg-rootfs 82ca1e17-8bc2-4895-96bb-3c39e5863238 /
vdb
vdc        
```

* Format volume partition
```
sudo mkfs.ext4 /dev/vdb
sudo mkfs.ext4 /dev/vdc
```

* Add mount point "assets" to `/home/deploy/assets`
```
sudo mkdir -p /home/deploy/assets
sudo chmod 777 /home/deploy/assets
sudo chown deploy:deploy /home/deploy/assets
```

* Add mount point "database" to `/home/deploy/database`
```
sudo mkdir -p /home/deploy/database
sudo chmod 777 /home/deploy/database
chown deploy:deploy /home/deploy/database
```

* Show mount point UUID
```
sudo blkid
-------
/dev/vda1: LABEL="cloudimg-rootfs" UUID="82ca1e17-8bc2-4895-96bb-3c39e5863238" TYPE="ext4" PARTUUID="3276fc6f-01"
/dev/vdb: UUID="12c75a16-c392-4dd6-b031-d5acd09fb75f" TYPE="ext4"
/dev/vdc: UUID="ff0301b7-1198-4b3d-afd4-708044cf38d8" TYPE="ext4"
```

* Edit `/etc/fstab` file
```
LABEL=cloudimg-rootfs                         /                      ext4     defaults    0 0
UUID=12c75a16-c392-4dd6-b031-d5acd09fb75f     /home/deploy/assets    ext4     defaults    0 0
UUID=ff0301b7-1198-4b3d-afd4-708044cf38d8     /home/deploy/database  ext4     defaults    0 0
```

* Remounting all partition
```
sudo mount -a
```

* View all partition mounted
```
sudo mount -v
-------
/dev/vdc on /home/deploy/database type ext4 (rw,relatime,data=ordered)
/dev/vdb on /home/deploy/assets type ext4 (rw,relatime,data=ordered)
lxcfs on /var/lib/lxcfs type fuse.lxcfs (rw,nosuid,nodev,relatime,user_id=0,group_id=0,allow_other)
/dev/vda1 on /var/lib/docker/plugins type ext4 (rw,relatime,data=ordered)
/dev/vda1 on /var/lib/docker/overlay2 type ext4 (rw,relatime,data=ordered)
tmpfs on /run/user/1001 type tmpfs (rw,nosuid,nodev,relatime,size=1643188k,mode=700,uid=1001,gid=1001)
```

* View mounting `/dev/vda` only
```
sudo mount -v | grep vda
```

* View mounting `/dev/vdb` only
```
sudo mount -v | grep vdb
```

* View mounting `/dev/vdc` only
```
sudo mount -v | grep vdc
```
