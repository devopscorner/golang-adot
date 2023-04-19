## Setup Account

Setup Account for VPS (Virtual Private Server) Ubuntu version 16.04

### Create Deploy User & Authorize
```
sudo adduser deploy            # password: d3pl0y_zeroc0d3 (default)
sudo usermod -aG sudo ubuntu
sudo usermod -aG sudo ec2-user
sudo usermod -aG sudo deploy

sudo visudo
-----
root      ALL=(ALL:ALL) ALL
ubuntu    ALL=(ALL:ALL) ALL
deploy    ALL=(ALL:ALL) ALL
```

### Generate Public-Key (**without password**)
```
/usr/bin/ssh-keygen -t rsa -b 4096 -C "deploy@devopscorner.cloud9" -f $HOME/.ssh/id_rsa -q -N "";
```

### Generate `pem` file
```
/usr/bin/openssl rsa -in ~/.ssh/id_rsa -outform pem > id_rsa.pem
```
