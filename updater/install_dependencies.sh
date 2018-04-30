#!/bin/bash
dir=`dirname $0`
while read p; do
    go get -u -v ${p}
done < ${dir}/dependencies.txt