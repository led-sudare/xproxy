#!/bin/sh
cname=`cat ./cname`
docker build ./ -t $cname 
docker run -t --init --name $cname -v `pwd`:/work/ -p 5510:5510/tcp -p 5511:5511/tcp $cname 