#!/usr/bin/env bash

set -o errtrace
set -o errexit

CRANE_TAR_URL=http://ocrqkagax.bkt.clouddn.com/crane.tar.gz

# setup latest release tag
CRANE_RELEASE=$1
if [ -z $CRANE_RELEASE ]
then
  CRANE_RELEASE=v1.0.4
fi

# make sure curl command exists
if ! command -v curl
then
  echo "curl command not found, install curl please "
  echo ""
  echo "1, yum install curl"
  echo "2, apt-get install curl"
  exit 1
fi

# download
curl -sSL  ${CRANE_TAR_URL} | tar xvzf -

echo "Enter IP address that your want bind Crane service [ENTER]"
read listener_ip

cd crane && CRANE_IP=${listener_ip} VERSION=${CRANE_RELEASE} ./deploy.sh
cd -
