#!/bin/bash

set -o errtrace
set -o errexit

docker run --rm -v $(pwd)/frontend:/data digitallyseamless/nodejs-bower-grunt:5 bower install
docker run --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make

if [ ! -f docker/docker ]; then
    curl https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz | tar xzv
fi

export TAG=`git log --pretty=format:'%h' -n 1`
export CRANE_SWARM_MANAGER_IP=$CRANE_IP
export REGISTRY_PREFIX=""

# node env check
echo "Checking the node status"
./frontend/misc-tools/node-init.sh

# swarm init
echo "Trying to init swarm cluster"
INIT_ERROR=$(docker swarm init --advertise-addr=$CRANE_IP 2>&1 > /dev/null) || {
   docker info 2>/dev/null | grep Swarm | grep -v inactive || {
      printf "\033[41mERROR:\033[0m failed to init swarm against cmd: \e[1;34mdocker swarm init --advertise-addr=$CRANE_IP\e[0m\n"
      echo "$INIT_ERROR"
      exit 1
   }
}
echo "Swarm cluster have been running!"

docker-compose -p crane -f deploy/docker-compose.yml stop
docker-compose -p crane -f deploy/docker-compose.yml rm -f
docker-compose -p crane -f deploy/docker-compose.yml up -d
