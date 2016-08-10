#!/bin/sh

set -e

sed -i "s/ROLEX_IP/$ROLEX_IP/;s/ROLEX_PORT/$ROLEX_PORT/" /etc/docker/registry/config.yml
/bin/registry "$@"
