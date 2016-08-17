#!/bin/bash

# test help
assert_status_code "get" "misc/v1/config" 200

# test feature enabled
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep account &>/dev/null
if [  "$?" == "0" ] 
then
  ok "feature account enabled"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep registry &>/dev/null
if [  "$?" == "0" ] 
then
  ok "feature registry enabled"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep logging &>/dev/null
if [  "$?" == "0" ] 
then
  ok "feature logging enabled"
fi
