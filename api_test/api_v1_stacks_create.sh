#!/bin/bash

# create stack
msg "create a stack from bundle with invalid json content"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/api/v1/stacks Authorization:$token &>/dev/null foo=bar
if [ "$?" != "5" ]
then
  fail "create stack error"
else
  ok "create stack with invalid json response header should be 503"
fi

msg "create a stack from bundle"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/api/v1/stacks Authorization:$token foo=bar | jq .code | grep 11502 1>/dev/null 2>&1
if [ "$?" != "0" ]
then
  fail "create stack error"
else
  ok "create stack with invalid json error code is 11502"
fi

stack_name=`cat data_stack_create_correct.json | jq .Namespace | tr -d '"'`

#msg "create a stack from wrong data"
#http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/api/v1/stacks?group_id=$group_id @data_stack_create_incorrect.json Authorization:$token

#http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/api/v1/stacks?group_id=$group_id @data_stack_create_incorrect.json Authorization:$token | jq .data | grep "success" 1>/dev/null 2>&1
#if [ "$?" != "4" ]
#then
  #fail "create stack error"
#else
  #ok "fail to create a stack $stack_name"
#fi

msg "create a stack from bundle"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/api/v1/stacks?group_id=$group_id @data_stack_create_correct.json Authorization:$token | jq .data | grep "success" 1>/dev/null 2>&1
if [ "$?" != "0" ]
then
  fail "create stack error"
else
  ok "successfully create a stack $stack_name"
fi

