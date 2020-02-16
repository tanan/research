#!/bin/bash

docker rm -f research-db
docker run -p 5432:5432 -d --name "research-db" research-db:latest
