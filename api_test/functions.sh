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

function assert_status_code () {
  msg "http $1 $2"
  status_code="$(http -h --timeout=4.5 $1 $SERVER_PATH/$2 Authorization:$token | grep HTTP/  | cut -d ' ' -f 2)"
  if [ "$status_code" != "$3" ]
  then
    fail "request $2 $1 respone status code get $? not $3"
  else
    ok "request $2 $1 got response status code $3"
  fi
}

