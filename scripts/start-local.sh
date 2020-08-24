#!/usr/bin/env bash

set -e

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"
ROOT_DIR="$(git rev-parse --show-toplevel)"

echo "executing ${SCRIPT_NAME}"

cd ${ROOT_DIR}

## Check if gogorup is installed
if ! ttab_loc="$(type -p ttab)" || [[ -z ${ttab_loc} ]]; then
  echo "ttab is not installed. installing...."
  sudo npm install ttab -g
fi

## declare list of all services
declare -a services=("${SCRIPT_DIR}/run-btc-wallet.sh"
)

## now loop through the above array
for svc in "${services[@]}"; do
  echo "running $svc"
  ttab -a iTerm2 sh "${svc}"
done

echo "all services run"
