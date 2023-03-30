#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"

echo "${SCRIPT_NAME} is running... "

compose_cmd="docker compose"

if ! command -v docker
then
 printf "Cannot check docker, please install docker:
        https://docs.docker.com/get-docker/ \n"
   exit 1
fi

${compose_cmd} -p btcwallet_dev \
  -f "${REPO_ROOT}/deployments/docker-compose/dev/docker-compose.yml" \
  up --detach --build

echo "${SCRIPT_NAME} done."
