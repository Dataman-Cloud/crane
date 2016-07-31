#!/bin/bash

docker run -it --rm -w /go/src/github.com/Dataman-Cloud/rolex -v $(pwd):/go/src/github.com/Dataman-Cloud/rolex golang:1.5.4 make
docker run -it --rm -w /go/src/github.com/Dataman-Cloud/rolex -v $(pwd):/go/src/github.com/Dataman-Cloud/rolex golang:1.5.4 make collect-cover-data
docker run -it --rm -w /go/src/github.com/Dataman-Cloud/rolex -v $(pwd):/go/src/github.com/Dataman-Cloud/rolex golang:1.5.4 make test-cover-html
docker run -it --rm -w /go/src/github.com/Dataman-Cloud/rolex -v $(pwd):/go/src/github.com/Dataman-Cloud/rolex golang:1.5.4 make test-cover-func
