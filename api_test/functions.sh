#!/bin/bash

function fail () {
  echo  -e " \033[;31m ~ [FAIL] $1 \033[0m "
  exit 1
}

function ok () {
  echo  -e " \033[;32m ~ [OK] $1 \033[0m "
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
  else
    ok "request $1 got response header 200"
  fi
}

function assert_400 () {
  msg "http $1 $2"
  http --check-status --ignore-stdin --timeout=4.5 $1 $SERVER_PATH/$2 &>/dev/null
  if [ "$?" != "4" ]
  then
    fail "request $1 didn't get response $?"
  else
    ok "request $1 got response header 400"
  fi
}

function assert_404 () {
  msg "http $1 $2"
  http --check-status --ignore-stdin --timeout=4.5 $1 $SERVER_PATH/$2 &>/dev/null
  if [ "$?" != "4" ]
  then
    fail "request $1 didn't get response $?"
  else
    ok "request $1 got response header 404"
  fi
}

