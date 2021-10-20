#!/bin/sh

docker run --rm -it -v $(pwd):/app -p 8980:80 --workdir="/app" golang go run main.go
