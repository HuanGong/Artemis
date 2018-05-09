#!/bin/bash

curDir=`pwd`
baseDir=$(cd "$(dirname "$0")"; pwd)

GO_PATH=${GOPATH%%:*}


# pushd $baseDir
cd $baseDir
echo `pwd`

protoc ./*.proto --go_out=plugins=grpc:.