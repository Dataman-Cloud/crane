#!/bin/bash

export ROLEX_SWARM_MANAGER_IP=$ROLEX_IP
export TAG=${VERSION:-1.0}
export DOCKER_COMPOSE=${VERSION:-.}/docker-compose.yml
export NODE_INIT=${VERSION:-../misc-tools}/node-init.sh
export REGISTRY_PREFIX=${REGISTRY_PREFIX:-catalog.shurenyun.com/library/}

# node env check
echo "Checking the node status"
$NODE_INIT || exit 1

# swarm init
echo "Trying to init swarm cluster"
$(docker swarm init --advertise-addr=$ROLEX_IP &>/dev/null) || {
   echo "Swarm cluster have been running!"
}

docker-compose -p rolex -f $DOCKER_COMPOSE up -d

# feedback the activities
curl -XPOST 123.59.58.58:4500/activities -H "Content-Type: application/json" -d'{"UniqId": "'"$(hostname)"'"}' &>/dev/null || exit 1
