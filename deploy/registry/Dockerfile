FROM registry:2.5.1

MAINTAINER weitao zhou <wtzhou@dataman-inc.com>

ENV CRANE_PORT 5013
ADD ./config.yml.template /etc/docker/registry/config.yml
ADD ./private_key.pem /etc/registry/private_key.pem
ADD ./root.crt /etc/registry/root.crt

COPY docker-entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["/etc/docker/registry/config.yml"]
