#!/bin/bash

X=`docker images | grep none | awk '{ print $3}'`
for i in $X; do   docker rmi $X; done
