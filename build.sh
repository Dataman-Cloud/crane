#!/bin/bash

REGISTRY_PREFIX=${REGISTRY_PREFIX:-demoregistry.dataman-inc.com/library/}
VERSION=${VERSION:-v1.0.0rc5}

docker run -it --rm -v $(pwd):/data digitallyseamless/nodejs-bower-grunt:5 bower install

docker build -f ./deploy/Dockerfile -t ${REGISTRY_PREFIX}blackmamba:${VERSION} .
