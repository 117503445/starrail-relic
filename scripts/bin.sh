#!/usr/bin/env bash

# should be run in Dockerfile.builder Environment

set -e

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./