#!/usr/bin/env bash

set -e

migrate -database 'mysql://root@tcp(127.0.0.1:3306)/btc_wallet' -path internal/db/migrations down

echo "Done."