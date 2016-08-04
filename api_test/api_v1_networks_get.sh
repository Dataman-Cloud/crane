#!/bin/bash


first_network=`http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/networks Authorization:$token | jq ".data[0].Id" | tr -d '"'`
if [ -z "$first_network" ]
then
  fail "got the first network fail"
else
  ok "got the first network $first_network"
fi


# inspect network
msg "inspect network"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/networks/$first_network Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "inspect networks failed"
else
  ok "inspect network success"
fi
