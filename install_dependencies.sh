#!/bin/bash
while read p; do
    go get -v ${p}
done < dependencies.txt