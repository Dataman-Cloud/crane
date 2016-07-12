#! /usr/bin/env bash

docker run -it --rm -v $(pwd):/data digitallyseamless/nodejs-bower-grunt:5 bower install
docker run --net host --name BlackMamba -v $(pwd):/usr/share/nginx/html:ro -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx:stable-alpine
