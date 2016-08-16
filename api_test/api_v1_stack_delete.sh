#!/bin/bash


# delete a stack
msg "delete a stack thr api"
http --check-status --ignore-stdin --timeout=4.5 delete $SERVER_PATH/api/v1/stacks/$stack_name  Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "delete stack failed"
else
  ok "delete stack success"
fi

# delete a stack
msg "delete a noexists stack thr api"
http --check-status --ignore-stdin --timeout=4.5 delete $SERVER_PATH/api/v1/stacks/noexists  Authorization:$token &>/dev/null
if [ "$?" != "4" ]
then
  fail "delete stack failed"
else
  ok "delete stack with namespace not exists should return 404"
fi
