#!/bin/bash
export GOPATH=`pwd`
mkdir -p bin
export GOBIN=`pwd`/bin
go get -t ./...
go build .

