#!/bin/sh

set -e

sed -i "s/CRANE_IP/$CRANE_IP/;s/CRANE_PORT/$CRANE_PORT/" /etc/docker/registry/config.yml

case "$1" in
  *.yaml|*.yml) set -- registry serve "$@" ;;
  serve|garbage-collect|help|-*) set -- registry "$@" ;;
esac

exec "$@"
