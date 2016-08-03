#!/bin/bash


# list networks
msg "list networks"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/networks Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "list networks failed"
fi
