#!/bin/sh

set -e

sed -i "s/CRANE_IP/$CRANE_IP/;s/CRANE_PORT/$CRANE_PORT/" /etc/docker/registry/config.yml
/bin/registry "$@"
