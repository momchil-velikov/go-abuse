#! /bin/sh

#GO=$HOME/src/go/src/cmd/go/go
GO=go

GOPATH=$(pwd)/p1 $GO install a
GOPATH=$(pwd)/p2 $GO install a
GOPATH=$(pwd)/p2:$(pwd) $GO install c
GOPATH=$(pwd)/p1:$(pwd) $GO build b
