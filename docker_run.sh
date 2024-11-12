#!/bin/sh

mkdir ./dist
rm -rf ./dist/*

docker run --rm -it -v $(pwd):/app -e CGO_ENABLED=0 --workdir="/app" golang sh -c '\
  for os in windows linux darwin; do \
    for arch in amd64 arm64; do \
      env GOOS=$os GOARCH=$arch go build -o "./dist/main_${os}_${arch}" main.go || exit 1; \
    done; \
  done'