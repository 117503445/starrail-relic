#!/usr/bin/env bash

# should be run in Dockerfile.builder Environment

set -e

cd /workspace/scripts/builder
go run main.go
cd /workspace

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./