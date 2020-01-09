#!/usr/bin/env bash

echo "Building patch proto files"
protoc --go_out=. patch.proto