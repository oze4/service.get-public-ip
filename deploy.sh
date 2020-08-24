#!/bin/bash

docker build \
  --pull \
  --rm -f "Dockerfile" \
  -t oze4/service.get-public-ip:latest "." \
&& \
docker push oze4/service.get-public-ip:latest