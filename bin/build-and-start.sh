#!/bin/bash

docker run -it --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make
curl https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz | tar xzv

export TAG=`git log --pretty=format:'%h' -n 1`
export CRANE_SWARM_MANAGER_IP=$CRANE_IP
export REGISTRY_PREFIX=""

=======
export TAG=1.0
export CRANE_SWARM_MANAGER_IP=$CRANE_IP
export REGISTRY_PREFIX=demoregistry.dataman-inc.com/library/

docker-compose -p crane -f deploy/docker-compose.yml stop
docker-compose -p crane -f deploy/docker-compose.yml rm -f
docker-compose -p crane -f deploy/docker-compose.yml up -d
