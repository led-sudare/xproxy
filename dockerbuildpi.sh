#!/bin/sh
cname=`cat ./cname`
docker build ./ -t $cname 

echo "stopping.. " && docker container stop $cname
echo "removing.. " && docker container rm $cname

echo "run and stating.. " && docker run -d --init --name $cname --net=host -v `pwd`:/work/ -p 5510:5510/tcp -p 5511:5511/tcp --restart=always $cname 
