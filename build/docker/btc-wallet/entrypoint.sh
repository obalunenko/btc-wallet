#!/bin/sh

set -e

echo "current user $(whoami)"

./btc-wallet \
  --input="${btcwallet_INPUT}" \
  --result="${btcwallet_RESULT}" \
  --archive="${btcwallet_ARCHIVE}" \
  --errors="${btcwallet_RECEIVE_ERRORS}" \
  --log_level="${btcwallet_LOG_LEVEL}" \
  --log_format="${btcwallet_LOG_FORMAT}" \
  --log_sentry_dsn="${btcwallet_LOG_SENTRY_DSN}" \
  --log_sentry_trace="${btcwallet_LOG_SENTRY_TRACE}" \
  --log_sentry_trace_level="${btcwallet_LOG_SENTRY_TRACE_LEVEL}"
