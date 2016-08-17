#!/bin/bash

# test help
assert_status_code "get" "misc/v1/help" 200
