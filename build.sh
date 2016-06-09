#!/bin/bash
export GOPATH=`pwd`
mkdir bin
export GOBIN=`pwd`/bin
go get ./...
go build .

