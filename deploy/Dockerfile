FROM alpine:3.4

MAINTAINER weitao zhou <wtzhou@dataman-inc.com>

ADD ./docker/docker /usr/bin/docker
ADD ./bin/rolex /go/bin/rolex
ADD ./deploy/env /go/bin/env_file
ADD ./doc/registry/private_key.pem /go/bin/private_key.pem
ADD ./bin/catalog /go/bin/catalog
ADD ./img /go/bin/img
WORKDIR /go/bin

ENTRYPOINT ./rolex
