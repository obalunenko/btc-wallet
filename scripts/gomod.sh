#!/usr/bin/env sh

set -e

go mod tidy -v
go mod vendor
go mod verify

echo "Done."