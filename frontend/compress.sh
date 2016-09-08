#!/bin/sh

npm install -g gulp && npm install

rm -rf ./dist/*

bower install

gulp
