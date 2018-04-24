#!/bin/bash

binFile=uolo_centre_linux

killall $binFile

sleep 5

cp new $binFile

nohup ./$binFile >> ./log.log 2>&1 &