#!/bin/bash

if [ $# -ne 2 ]; then
  echo "usage: $0 machine config"
  exit -1
fi

machine=$1
configFile=$2

#dataFile=$3
servicePath=/data/deploy/uolo/api_center
binFile=uolo_center_linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $binFile main.go

if [ $? -ne 0 ]; then
  echo "build failed"
  exit -1
fi

ssh $machine "mkdir -p $servicePath;"
scp ./$binFile    $machine:$servicePath/new
#scp ./$configFile $machine:$servicePath/config.toml

scp ./restart.sh $machine:$servicePath
#scp -r ./$dataFile $machine:$servicePath

#ssh $machine "cd $servicePath; pwd ; /bin/bash restart.sh"

if [ $? -ne 0 ]; then
  echo "run failed"
  exit -1
fi

echo "run success"
