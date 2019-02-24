#!/usr/bin/env bash

# Kind of a clunky way of ensuring we are in the proper directory but it
# handles a couple of use cases which have bit me in the past.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null 2>&1 && pwd )"
pushd $DIR/../cmd/server

# location for our binaries making sure to remove it when we are done
mkdir tmp

# Compression
# - The s & w (linker) flags strip debugging info and reduce size by ~25%.
# - Note: you tried to get some additional compreesion via `upx` and while it
#   reduced the size quite nicely it also resulted in failures. You can try
#   agiain later but for now, the size difference is negligible.
GOOS=linux go build -ldflags="-s -w" -o tmp/main main.go

pushd tmp
pwd
zip main.zip main

# You didn't want to configure the region due to other projects so that is why I
# added it here:
aws lambda update-function-code \
  --function-name tic-tac-toe-move \
  --zip-file fileb://main.zip \
  --region us-east-1

popd
rm -rf ./tmp
