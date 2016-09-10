#!/bin/bash

docker run -it --rm -w /go/src/github.com/Dataman-Cloud/crane -v $(pwd):/go/src/github.com/Dataman-Cloud/crane golang:1.5.4 make
curl https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz | tar xzv

export TAG=1.0
export CRANE_SWARM_MANAGER_IP=$CRANE_IP
export REGISTRY_PREFIX=demoregistry.dataman-inc.com/library/
docker-compose -p crane -f deploy/docker-compose.yml stop
docker-compose -p crane -f deploy/docker-compose.yml rm -f

# remove the deprecated image: crane:$TAG , and triger build action again
# have to rm the image specially to avoid mysql/registry build.
docker rmi -f ${REGISTRY_PREFIX}crane:$TAG

docker-compose -p crane -f deploy/docker-compose.yml up -d
