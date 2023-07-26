#!/bin/bash

sed -i "/#\$nrconf{restart} = 'i';/s/.*/\$nrconf{restart} = 'a';/" /etc/needrestart/needrestart.conf

apt update
apt upgrade -y
apt install -y golang

mkdir -p bin

echo "#!/bin/bash" > bin/api
echo "source vars/database.env" >> bin/api
echo "cd api && git stash && git pull && go run ." >> bin/api
chmod a+rx bin/api