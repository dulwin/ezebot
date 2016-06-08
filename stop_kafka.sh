#! /bin/bash

docker-compose kill zookeeper kafka;
docker-compose rm -f zookeeper kafka;
