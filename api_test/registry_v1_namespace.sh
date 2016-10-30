#!/bin/bash

namespace=admintest
# create namespace
msg "create a registry namespace for user admin@admin.com"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/registry/v1/namespace Authorization:$token namespace=$namespace &>/dev/null
if [ "$?" != "0" ]
then
  fail "failed to create namespace for user admin@admin.com"
fi

# get namespace
msg "get a registry namespace for user admin@admin.com"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/registry/v1/namespace Authorization:$token | jq .data.Namespace | grep "$namespace" &>/dev/null
if [ "$?" != "0" ]
then
  fail "failed to get namespace for user admin@admin.com"
fi
