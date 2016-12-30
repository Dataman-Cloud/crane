#!/bin/bash

docker pull  registry:2.5.1

docker run -d  --name registry -p 5000:5000  -p 5001:5001 -v `pwd`:/etc/registry/ -v $(pwd)/storage:/storage \
  -v `pwd`/config.yml.template:/etc/registry/config.yml registry:2.5.1 /etc/registry/config.yml
