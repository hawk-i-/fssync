#!/bin/bash

[[ -d ./.out ]] && rm -rf ./.out
mkdir -p ./.out/linux/amd64
# mkdir -p ./.out/windows/amd64

VERSION="dev"
TAG_LIST=$(git tag --points-at HEAD)
COMMIT_TAG=$(echo $TAG_LIST | cut -d ' ' -f 1)
COMMIT_HASH=$(git rev-parse --short HEAD)

if [ -n "$COMMIT_TAG" ]; then
    VERSION=$COMMIT_TAG

elif [ -n "$COMMIT_HASH" ]; then
    VERSION=$COMMIT_HASH
fi

GOARCH=amd64 GOOS=linux \
go build -ldflags "-s -w -X 'main.VERSION=$VERSION'" \
-trimpath \
-o ./.out/linux/amd64/fssync

# GOARCH=amd64 GOOS=windows \
# go build -ldflags "-s -w -X 'main.VERSION=$VERSION'" \
# -trimpath \
# -o ./.out/windows/amd64/fssync.exe