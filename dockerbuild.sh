#!/bin/sh
cname=`cat ./cname`
docker build ./ -t $cname --build-arg cname=$cname
docker run -t --init --name $cname -v `pwd`:/go/src/$cname/ -p 5510:5510/tcp -p 5511:5511/tcp $cname 