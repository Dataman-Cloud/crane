#!/bin/sh

set -o errtrace
set -o errexit

npm install -g gulp && npm install

rm -rf ./dist/*

bower install

gulp
