#!/bin/bash


# test login
msg "test login"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/account/v1/login password=adminadmin email=admin@admin.com &>/dev/null
if [ "$?" != "0" ]
then
  fail "login with email admin@admin.com and password admiadmin failed"
fi

# get token
msg "login token"
token=`http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/account/v1/login password=adminadmin email=admin@admin.com | jq .data | tr -d '"'`
msg "login got $token"
if [ -z "$token" ]
then
  fail "get token failed"
fi

