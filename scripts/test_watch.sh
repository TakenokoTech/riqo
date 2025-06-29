#!/usr/bin/env bash
set -euo pipefail
# Ensure reflex is installed
if ! command -v reflex &> /dev/null
then
    echo "reflex could not be found. Install it with: go install github.com/cespare/reflex@latest"
    exit 1
fi

# Run reflex to watch for file changes and execute tests
reflex -r '(\.go$|go\.mod)' -s -- sh -c 'go test ./...'
