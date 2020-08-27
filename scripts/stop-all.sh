#!/usr/bin/env bash
set -e

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(git rev-parse --show-toplevel)"

echo "executing ${SCRIPT_NAME}"

cd ${ROOT_DIR}

declare -a services=("btc-wallet"
)

## now loop through the above array
for svc in "${services[@]}"; do
  echo "killing ${svc}..."
  killall "${svc}"
done

echo "all services stopped"
