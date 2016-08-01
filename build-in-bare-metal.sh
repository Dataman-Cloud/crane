#!/bin/bash

# docker run --name rolex_db --net=host -e MYSQL_ROOT_PASSWORD=111111 -d mysql:5.7
docker run -it --rm -w /go/src/github.com/Dataman-Cloud/rolex -v $(pwd):/go/src/github.com/Dataman-Cloud/rolex golang:1.5.4 make
