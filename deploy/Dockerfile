FROM alpine:3.4

MAINTAINER weitao zhou <wtzhou@dataman-inc.com>

ADD ./bin/crane /go/bin/crane
ADD ./deploy/registry/private_key.pem /go/bin/private_key.pem
WORKDIR /go/bin
EXPOSE 5013

ENTRYPOINT ./crane
