#!/bin/bash

# create stack
msg "create a stack from bundle"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/api/v1/stacks Authorization:$token &>/dev/null foobar
if [ "$?" != "0" ]
then
  fail "list nodes failed"
fi
