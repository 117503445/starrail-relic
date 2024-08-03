#!/usr/bin/env bash

# build bin By Docker

set -e

docker build -t starrail-relic-builder -f Dockerfile.builder .
docker run --rm -v $PWD:/workspace starrail-relic-builder