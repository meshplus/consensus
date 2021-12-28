#!/bin/bash


export GOROOT=$(go env GOROOT)
export GOPATH=$(go env GOPATH)
mkdir -p $GOPATH/src/github.com/meshplus/
cd ..
echo "[[[[" `pwd` "]]]]"
mv consensus $GOPATH/src/github.com/meshplus/
cd $GOPATH/src/github.com/meshplus/consensus
