#!/bin/bash

# This script builds the hakjctl binary.
# Output directory: ./bin/hakjctl/
# Execute this script from the project root.

go build -o ./bin/hakjctl/ ./cmd/hakjctl/
