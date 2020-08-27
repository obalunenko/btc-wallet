#!/usr/bin/env sh
set -e
echo "Building..."

BIN_OUT=./bin/btc-wallet

go build -o ${BIN_OUT} ./cmd/btc-wallet

echo "Binary compiled at ${BIN_OUT}"