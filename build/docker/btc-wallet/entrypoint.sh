#!/usr/bin/env sh
set -e

LOGNAME=${APP_NAME}
LOGDIR=${ROOT_DIR}/logs
mkdir -p logs

LOGFILE=${LOGDIR}/${LOGNAME}__$(date "+%d-%m-%Y_%H-%M-%S").log


./btc-wallet \
  --db_conn_max_lifetime=${BTCWALLET_DB_CONN_MAX_LIFE} \
  --db_max_idle_conns=${BTCWALLET_DB_MAX_IDLE_CONNS} \
  --db_max_open_conns=${BTCWALLET_DB_MAX_OPEN_CONS} \
  --log_level=${BTCWALLET_LOG_LEVEL} \
  --log_file=${BTCWALLET_LOG_FILE} \
  --db=${BTCWALLET_DB}