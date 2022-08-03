#!/bin/bash
git pull
CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o redis-manager main.go

docker build -f Dockerfile -t redis-manager . 
redis-manager=`docker ps -a | grep redis-manager |wc -l`

if [ ${redis-manager} == 1 ]
   then
      id=`docker ps -a| grep redis-manager| awk '{print $1}'`
      echo "docker stop $id"
      docker stop $id
      echo "docker rm $id"
      docker rm $id
      echo "docker run redis-manager"
      docker run -d -it -v /data/apps/redis-manager/logs/:/data/logs  -p  8000:8000 --restart=always  --name redis-manager redis-manager
else
   echo "docker run redis-manager"
   docker run -d -it  -v /data/apps/redis-manager/logs/:/data/logs  -p  8000:8000 --restart=always  --name redis-manager redis-manager
fi
