#!/bin/bash

# test login
msg "test login"
http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/account/v1/login password=adminadmin email=admin@admin.com &>/dev/null
if [ "$?" != "0" ]
then
  fail "login with email admin@admin.com and password admiadmin failed"
else
  ok "login success with admin@admin.com / adminadmin"
fi

# test login with invaled password
msg "test login with invalid password -- httpCode"
#verify httpcode
httpCode=`http --check-status --ignore-stdin --timeout=4.5 -h post http://192.168.1.102/account/v1/login password=adminadmin email=admin@admin.co 2>/dev/null |awk -F ' ' '/HTTP/{print $2}'`
if [ $httpCode != "503" ]
then
  fail "login with invalid account，httpCode isn't correct."
else
  ok "login with invalid account，httpCode is correct."
fi

#verify code
msg "test login with invalid password -- code"
code=`http --check-status --ignore-stdin --timeout=4.5 post http://192.168.1.102/account/v1/login password=adminadmin email=admin@admin.co 2>/dev/null |awk -F ':|,' '/code/{print $2}'`
if [ $code != "12007" ]
then
  fail "login with invalid account，code isn't correct."
else
  ok "login with invalid account，code is correct."
fi

# get token
msg "login token"
token=`http --check-status --ignore-stdin --timeout=4.5 post $SERVER_PATH/account/v1/login password=adminadmin email=admin@admin.com | jq .data | tr -d '"'`
if [ -z "$token" ]
then
  fail "get token failed"
else
  ok "got token $token"
fi

# get account
msg "my account"
account=`http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/account/v1/aboutme Authorization:$token`
account_email=`echo $account | jq .data.Email | tr -d '"'`
account_id=`echo $account | jq .data.Id | tr -d '"'`

if [ -z "account" ]
then
  fail "get account info failed"
else
  ok "got account info, email is `echo $account | jq .data.Email`"
fi

# get account groups
msg "get groups of a account"
groups=`http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/account/v1/accounts/$account_id/groups Authorization:$token`
group_id=`echo $groups | jq ".data [0].Id"`
group_name=`echo $groups | jq ".data [0].Name" | tr -d '"'`

if [ -z "groups" ]
then
  fail "get groups failed"
else
  ok "got group info, id is `echo $group_id` and name is `echo $group_name`"
fi
