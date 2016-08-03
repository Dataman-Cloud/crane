#!/bin/bash


first_network=`http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/networks Authorization:$token | jq ".data[0].Id" | tr -d '"'`
msg "got first $first_network"


# inspect network
msg "inspect network"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/networks/$first_network Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "inspect networks failed"
fi
