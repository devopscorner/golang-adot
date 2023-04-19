## Swap Configuration

Swap space configuration for Ubuntu 16.04

* Check available swap
```
sudo swapon --show
sudo free -h
```

* Check space available
```
sudo df -h
```

* Create swap file
```
sudo fallocate -l [size_in_giga] /swapfile
-------
sudo fallocate -l 8.0G /swapfile    # set 8GByte
ls -lh /swapfile                    # verify swap
```

* Enabling swap file
```
sudo chmod 600 /swapfile
ls -lh /swapfile                    # verify permission
```

* Marked swap file
```
sudo mkswap /swapfile
-------
Setting up swapspace version 1, size = 8 GiB (8589930496 bytes)
no label, UUID=267498be-04b8-48d0-bf67-e43f18d03e49
```

* Allow system to start
```
sudo swapon /swapfile
```

### Test & Validation
* Verify swap available
```
sudo swapon --show
-------
NAME      TYPE SIZE USED PRIO
/swapfile file   8G   0B   -1
```

* Check output free
```
free -h
```

### Setup `/etc/fstab` Swap 
* Backup old configuration `/etc/fstab`
```
sudo cp /etc/fstab /etc/fstab.bak
```

* Add swap file information
```
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

* Adjusting the Cache Pressure Setting
```
sudo sysctl vm.vfs_cache_pressure=50
sudo sysctl vm.swappiness=10
```

* Edit configration `/etc/sysctl.conf`, add in bottom line:
```
vm.swappiness=10
vm.vfs_cache_pressure=50
```

* Remounting all partition & swap
```
sudo mount -a
```

* View free memory
```
sudo free -m
-------
              total        used        free      shared  buff/cache   available
Mem:          16046         139       15152          20         753       15568
Swap:         32767           0       32767

sudo free -h
-------
Mem:            15G        141M         14G         20M        754M         15G
Swap:           31G          0B         31G
```

### Notes
* References:
[Digital Ocean - Swap](https://www.digitalocean.com/community/tutorials/how-to-add-swap-space-on-ubuntu-16-04)