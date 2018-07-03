#!/bin/bash

binFile="uolo_center_linux"
killall "$binFile"; sleep 5;
nohup ./"$binFile" >> ./log.log 2>&1 &