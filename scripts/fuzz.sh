#!/usr/bin/env bash

# This script is used to fuzz the package.
# It iterates through all packages, finds fuzz tests, and fuzzes each one.
# It ensures that the package is fuzzed and that it is working correctly.

set -euo pipefail

GO="${GO:-go}"
FUZZTIME="${FUZZTIME:-10s}"
export GOWORK="${GOWORK:-off}"

while IFS= read -r pkg; do
  while IFS= read -r name; do
    [ -n "$name" ] || continue
    echo "Fuzzing $pkg::$name"
    "$GO" test "$pkg" -run=^$ -fuzz="^${name}$" -fuzztime="$FUZZTIME"
  done < <("$GO" test -list '^Fuzz' "$pkg" | grep '^Fuzz' || true)
done < <("$GO" list ./...)
