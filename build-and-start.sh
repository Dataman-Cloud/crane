#!/bin/bash

docker run -it --rm -w /go/src/github.com/Dataman-Cloud/rolex -v $(pwd):/go/src/github.com/Dataman-Cloud/rolex golang:1.5.4 make
curl https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz | tar xzv

export TAG=1.0
export ROLEX_SWARM_MANAGER_IP=$ROLEX_IP
export REGISTRY_PREFIX=demoregistry.dataman-inc.com/library/
docker-compose -p rolex -f deploy/docker-compose.yml stop
docker-compose -p rolex -f deploy/docker-compose.yml rm -f

# remove the deprecated image: rolex:$TAG , and triger build action again
# have to rm the image specially to avoid mysql/registry build.
docker rmi -f ${REGISTRY_PREFIX}rolex:$TAG

docker-compose -p rolex -f deploy/docker-compose.yml up -d
