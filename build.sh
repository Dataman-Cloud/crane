#!/bin/bash

REGISTRY_PREFIX=${REGISTRY_PREFIX:-demoregistry.dataman-inc.com/library/}
VERSION=${VERSION:-v1.0.0rc5}

docker build -f ./deploy/compress.Dockerfile -t ${REGISTRY_PREFIX}compress:v1.0.0 .

docker run -it --rm -v $(pwd):/data ${REGISTRY_PREFIX}compress:v1.0.0

docker build -f ./deploy/Dockerfile -t ${REGISTRY_PREFIX}blackmamba:${VERSION} .
