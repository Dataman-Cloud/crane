#! /usr/bin/env bash

docker run -it --rm -v $(pwd):/data digitallyseamless/nodejs-bower-grunt:5 bower install
docker run --net host --add-host rolex:192.168.59.105 --name BlackMamba -v $(pwd):/usr/share/nginx/html:ro -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx:stable-alpine
