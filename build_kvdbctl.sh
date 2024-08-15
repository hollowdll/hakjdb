#!/bin/bash

# This script builds the kvdbctl binary.
# Output directory: ./bin/kvdbctl/

go build -o ./bin/kvdbctl/ ./cmd/kvdbctl/
