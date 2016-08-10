FROM alpine:3.4

MAINTAINER weitao zhou <wtzhou@dataman-inc.com>

ADD ./docker/docker /usr/bin/docker
ADD ./bin/rolex /go/bin/rolex
WORKDIR /go/bin

ENTRYPOINT ./rolex
