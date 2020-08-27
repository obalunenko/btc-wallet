#!/usr/bin/env bash

set -e

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
ROOT_DIR="$(git rev-parse --show-toplevel)"
SCRIPT_DIR=${ROOT_DIR}/scripts

DB_NAME='btc_wallet'

echo "executing ${SCRIPT_NAME}"


cd "${ROOT_DIR}" || exit

# First create and seed the console database:
echo "DROP DATABASE IF EXISTS ${DB_NAME};" | mysql -h 127.0.0.1 -uroot
echo "CREATE DATABASE ${DB_NAME};" | mysql -h 127.0.0.1 -uroot

echo "database ${DB_NAME} created"