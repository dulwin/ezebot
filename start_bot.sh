#! /bin/bash

docker-compose kill bot producer;
docker-compose up --force-recreate --no-deps bot producer; 
