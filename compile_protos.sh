#!/bin/bash

# This script generates Go source code from the proto files.

protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  api/**/*.proto
