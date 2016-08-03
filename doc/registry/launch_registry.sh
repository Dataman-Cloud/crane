#!/bin/bash

docker pull  demoregistry.dataman-inc.com/srypoc/registry:2.3.0
docker tag demoregistry.dataman-inc.com/srypoc/registry:2.3.0 registry

docker run -d  --name registry -p 5000:5000  -p 5001:5001 -v `pwd`:/etc/registry/ -v $(pwd)/storage:/storage registry /etc/registry/config.yml

