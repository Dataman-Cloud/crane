#!/bin/bash

SERVER_PATH=localhost:5013

if ! command -v http &>/dev/null ; then
  echo "httpie not installed, 'apt-get install httpie' or  'brew install httpie'"
  exit 1
fi


if ! command -v jq &>/dev/null ; then
  echo "httpie not installed, 'apt-get install jq' or  'brew install jq'"
  exit 1
fi


function fail () {
  echo "[FAIL] $1"
  exit 1
}

function ok () {
  echo "[OK] $1"
}

function msg () {
  echo "[INFO] $1"
}

function assert_200 () {
  msg "http $1 $2"
  http --check-status --ignore-stdin --timeout=4.5 $1 $SERVER_PATH/$2 &>/dev/null
  if [ "$?" != "0" ]
  then
    fail "request $1 didn't get response $?"
  fi
}

function assert_400 () {
  msg "http $1 $2"
  http --check-status --ignore-stdin --timeout=4.5 $1 $SERVER_PATH/$2 &>/dev/null
  if [ "$?" != "4" ]
  then
    fail "request $1 didn't get response $?"
  fi
}

function assert_404 () {
  msg "http $1 $2"
  http --check-status --ignore-stdin --timeout=4.5 $1 $SERVER_PATH/$2 &>/dev/null
  if [ "$?" != "4" ]
  then
    fail "request $1 didn't get response $?"
  fi
}

# test help
assert_200 "get" "misc/v1/help"

# test config
assert_200 "get" "misc/v1/config"

# test feature enabled
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep account &>/dev/null
if [  "$?" == "0" ] 
then
  msg "account enabled"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep registry &>/dev/null
if [  "$?" == "0" ] 
then
  msg "registry enabled"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/misc/v1/config | jq .data.FeatureFlags | grep logging &>/dev/null
if [  "$?" == "0" ] 
then
  msg "logging enabled"
fi

# test health
assert_200 "get" "misc/v1/health"

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

# list nodes
msg "list nodes"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "list nodes failed"
fi

first_node=`http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes Authorization:$token | jq ".data[0].ID" | tr -d '"'`
msg "got first $first_node"


# inspect node
msg "inspect nodes"
http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes/$first_node Authorization:$token &>/dev/null
if [ "$?" != "0" ]
then
  fail "inspect nodes failed"
fi

http --check-status --ignore-stdin --timeout=4.5 get $SERVER_PATH/api/v1/nodes/$first_node Authorization:$token
