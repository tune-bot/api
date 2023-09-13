#!/bin/bash

sed -i "/#\$nrconf{restart} = 'i';/s/.*/\$nrconf{restart} = 'a';/" /etc/needrestart/needrestart.conf

apt update
apt upgrade -y
apt install -y golang

mkdir -p bin

echo "#!/bin/bash" > bin/api
echo "source vars/database.env" >> bin/api
echo "bin/download -U" >> bin/api
echo "cd api && go run src/*" >> bin/api
chmod a+rx bin/api