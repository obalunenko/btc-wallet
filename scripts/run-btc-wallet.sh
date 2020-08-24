#!/usr/bin/env sh
set -e

APP_NAME=btc-wallet

echo "Running... ${APP_NAME}"

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR="$( cd ${SCRIPT_DIR} && git rev-parse --show-toplevel )"

BIN=${ROOT_DIR}/bin

LOGNAME=${APP_NAME}
LOGDIR=${ROOT_DIR}/logs
mkdir -p logs

LOGFILE=${LOGDIR}/${LOGNAME}__$(date "+%d-%m-%Y_%H-%M-%S").log


${BIN}/${APP_NAME} \
  --db='mysql://root@tcp(127.0.0.1:3306)/btc_wallet?' 2>&1 | tee "${LOGFILE}"