#!/bin/bash

# get stack
msg "retrive a stack thr api"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/stacks/$stack_name  Authorization:$token | jq .data.Namespace | grep "$stack_name" &>/dev/null
if [ "$?" != "0" ]
then
  fail "get stack failed"
else
  ok "get stack success"
fi
