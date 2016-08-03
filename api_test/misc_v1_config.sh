#!/bin/bash

# test help
assert_200 "get" "misc/v1/config"

# test feature enabled
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep account &>/dev/null
if [  "$?" == "0" ] 
then
  msg "account enabled"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep registry &>/dev/null
if [  "$?" == "0" ] 
then
  msg "registry enabled"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep logging &>/dev/null
if [  "$?" == "0" ] 
then
  msg "logging enabled"
fi
