#! /usr/bin/env bash

set -o errtrace
set -o errexit

docker run --rm -v $(pwd):/data digitallyseamless/nodejs-bower-grunt:5 bower install
docker run --net host --add-host crane:$CRANE_IP --name BlackMamba -v $(pwd):/usr/share/nginx/html:ro -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx:stable-alpine
