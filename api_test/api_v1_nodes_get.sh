#!/bin/bash


first_node=`http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes Authorization:$token | jq ".data[0].ID" | tr -d '"'`

if [ -z "$first_node" ]; then
  fail "get first node failed"
else
  ok "got first $first_node"
fi


# inspect node
msg "inspect nodes"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes/$first_node Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "inspect nodes failed"
else
  ok "inspect node  $first_node success"
fi

