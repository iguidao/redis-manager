#!/bin/bash
git pull
CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o redis-web-manager main.go

docker build -f Dockerfile -t redis-web-manager . 
redis-web-manager=`docker ps -a | grep redis-web-manager |wc -l`

if [ ${redis-web-manager} == 1 ]
   then
      id=`docker ps -a| grep redis-web-manager| awk '{print $1}'`
      echo "docker stop $id"
      docker stop $id
      echo "docker rm $id"
      docker rm $id
      echo "docker run redis-web-manager"
      docker run -d -it -v /data/apps/redis-web-manager/logs/:/data/logs  -p  8000:8000 --restart=always  --name redis-web-manager redis-web-manager
else
   echo "docker run redis-web-manager"
   docker run -d -it  -v /data/apps/redis-web-manager/logs/:/data/logs  -p  8000:8000 --restart=always  --name redis-web-manager redis-web-manager
fi
