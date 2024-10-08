#!/bin/sh

docker run --rm -it -v $(pwd):/app -e CGO_ENABLED=0 -e GOOS=darwin -e GOARCH=amd64 --workdir="/app" golang go build main.go
