#!/bin/bash

# list nodes
msg "list nodes"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "list nodes failed"
fi
