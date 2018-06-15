#!/bin/bash

if [ $# -ne 1 ]; then
  echo "usage: $0 machine"
  exit -1
fi

machine=$1

servicePath=/data/deploy/uolo/center
binFile=uolo_center_linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $binFile main.go

if [ $? -ne 0 ]; then
  echo "build failed"
  exit -1
fi

ssh $machine "mkdir -p $servicePath;"
scp ./$binFile    $machine:$servicePath/new
#scp ./$configFile $machine:$servicePath/config.toml

#scp ./restart.sh $machine:$servicePath

ssh $machine "cd $servicePath; pwd ; /bin/bash restart.sh"

if [ $? -ne 0 ]; then
  echo "run failed"
  exit -1
fi

echo "run success"
