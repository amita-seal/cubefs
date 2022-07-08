#!/bin/bash

source ./init.sh

# build blobstore and get consul kafka mongo
INIT

docker build -t "blobstore:v1" -f Dockerfile .
docker run blobstore:v1  -d
