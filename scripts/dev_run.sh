#!/bin/bash

# Build the Riqo binary
echo "Building the Riqo binary..."
go build -o riqo
if [ $? -ne 0 ]; then
  echo "Build failed. Exiting."
  exit 1
fi

# Run the specified Riqo command
if [ $# -eq 0 ]; then
  echo "Usage: ./scripts/dev_run.sh <command> [args...]"
  echo "Example: ./scripts/dev_run.sh help"
  exit 1
fi

echo "Running './riqo $@'..."
./riqo "$@"
