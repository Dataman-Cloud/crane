#!/bin/bash


# list networks
msg "list networks"
assert_status_code "GET" "api/v1/networks" 200
