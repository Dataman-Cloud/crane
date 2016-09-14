#!/bin/bash

set -o errtrace
set -o errexit

export CRANE_SWARM_MANAGER_IP=$CRANE_IP
export TAG=${VERSION:-1.0}
export REGISTRY_PREFIX=${REGISTRY_PREFIX:-catalog.shurenyun.com/library/}

# node env check
echo "Checking the node status"
./node-init.sh || exit 1

# swarm init
echo "Trying to init swarm cluster"
$(docker swarm init --advertise-addr=$CRANE_IP &>/dev/null) || {
   echo "Swarm cluster have been running!"
}

docker-compose -p crane up -d

# feedback the activities
curl -XPOST 123.59.58.58:4500/activities -H "Content-Type: application/json" -d'{"UniqId": "'"$(hostname)"'"}' &>/dev/null || exit 1
