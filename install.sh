#!/usr/bin/env bash

set -o errtrace
set -o errexit

CRANER_TAR_URL=http://ocrqkagax.bkt.clouddn.com/craner.tar.gz

# setup latest release tag
CRANER_RELEASE=$1
if [ -z $CRANER_RELEASE ]
then
  CRANER_RELEASE=v1.0.3
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
curl -sSL  ${CRANER_TAR_URL} | tar xvzf -

echo "Enter IP address that your want bind Craner service [ENTER]"
read listener_ip

cd craner && ROLEX_IP=${listener_ip} VERSION=${CRANER_RELEASE} ./deploy.sh
cd -
