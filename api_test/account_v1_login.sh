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
