#!/bin/bash

# list stack
msg "retrive list stack thr api"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/stacks  Authorization:$token | jq ".data [].Namespace" | grep "$stack_name" &>/dev/null
if [ "$?" != "0" ]
then
  fail "list stack failed"
else
  ok "list stack success"
fi
