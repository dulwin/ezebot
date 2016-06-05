#! /bin/bash

docker-compose kill zookeeper kafka;
docker-compose rm -f zookeeper kafka;
docker-compose create zookeeper;
docker-compose scale zookeeper=3
docker-compose create kafka;
docker-compose scale kafka=3
