#!/usr/bin/env sh

rm -rf dist && mkdir dist

mkdir dist/linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wvs_proxy .

mv wvs_proxy dist/linux/wvs_proxy

mkdir dist/macos

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o wvs_proxy .

mv wvs_proxy dist/macos/wvs_proxy
