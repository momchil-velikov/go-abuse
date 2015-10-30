#! /bin/sh

GOPATH=$(pwd)/p1 go install a
GOPATH=$(pwd)/p2 go install a
GOPATH=$(pwd)/p2:$(pwd) go install c
GOPATH=$(pwd)/p1:$(pwd) go build b
