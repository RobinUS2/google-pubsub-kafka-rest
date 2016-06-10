#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$DIR
mkdir -p bin
export GOBIN=$DIR/bin
go get -v -t ./...
cd $DIR
go build -v .

