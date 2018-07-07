#!/bin/bash

buildMode=
remoteMachine=
serviceName="uolo_center"
binFile=uolo_center_linux
deployPath="/data/deploy/uolo/center"

if [ $# -lt 1 ];then
  echo "        Usage: cmd [-m] [-d]"
  echo "        -m [release,debug]"
  echo "        -r [machine name deploy this service]"
  exit 1
fi

while getopts "m:r:" arg #选项后面的冒号表示该选项需要参数
do
  case $arg in
    m)
      buildMode=$OPTARG;;
    r)
      remoteMachine=$OPTARG;;
    ?)
      echo "unkonw argument"; exit 1;;
  esac
done

if [ "$buildMode" != "release" ]; then
    buildMode="debug"
fi

echo "Build $binFile as $buildMode"

basePath=$(cd `dirname $0`; pwd)
cd "$basePath"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "$binFile" main.go
if [ $? -ne 0 ]; then
  echo "build $binFile failed"
  exit -1
fi

if [ "$buildMode" == "debug" ]; then
    exit 0
fi

#  =============  package the dist for release  ==============
if [ ! -d "dist/" ];then
  mkdir dist
fi

cp restart.sh dist
cp "$binFile" dist
cp -r ./data  dist
cp config_prd.toml dist/config.toml

deployFile="$serviceName"_$(date "+%Y%m%d_%H%M%S").tar.gz

tar -czf "$deployFile" ./dist/*
echo "Done! $deployFile is ready for deploy to server"
sleep 1;

if [ "$remoteMachine" != "" ]; then
    scp "$deployFile" "$remoteMachine":"$deployPath"
    if [ $? -ne 0 ]; then
        echo "Upload Service package to $remoteMachine:$deployPath Failed, err:$?"
        exit -1
    fi
    ssh "$remoteMachine" "cd $deployPath; tar -xvf $deployFile --strip-components 2 ; /bin/bash restart.sh"
    echo "Deploy to $remoteMachine and restart it; please check service status manually"
fi

sleep 1; echo "Clean Up ......."; rm "$deployFile"