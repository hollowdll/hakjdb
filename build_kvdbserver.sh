#!/bin/bash

# This script builds the kvdbserver binary.
# Output directory: ./bin/kvdbserver/

go build -o ./bin/kvdbserver/ ./cmd/kvdbserver/
