#!/bin/bash

# This script builds the hakjserver binary.
# Output directory: ./bin/hakjserver/
# Execute this script from the project root.

go build -o ./bin/hakjserver/ ./cmd/hakjserver/
