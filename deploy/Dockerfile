FROM alpine:3.4

MAINTAINER weitao zhou <wtzhou@dataman-inc.com>

ADD ./docker/docker /usr/bin/docker
ADD ./bin/crane /go/bin/crane
ADD ./deploy/env /go/bin/env_file
ADD ./doc/registry/private_key.pem /go/bin/private_key.pem
WORKDIR /go/bin

ENTRYPOINT ./crane
