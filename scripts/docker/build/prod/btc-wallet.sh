#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

SHORTCOMMIT="$(git rev-parse --short HEAD)"
VERSION=${VERSION:-""}

if [ -z "${VERSION}" ] || [ "${VERSION}" = "${SHORTCOMMIT}" ]
 then
  VERSION="v0.0.0-${SHORTCOMMIT}"
fi

DOCKER_REPO="${DOCKER_REPO:-"ghcr.io/obalunenko/btc-wallet"}"

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

checkInstalled 'docker'

docker build --rm --no-cache \
    -t "${DOCKER_REPO}btc-wallet:${VERSION}" \
    -t "${DOCKER_REPO}btc-wallet:latest" \
    -f "${REPO_ROOT}/build/docker/btc-wallet/Dockerfile" \
    --build-arg DOCKER_REPO="${DOCKER_REPO}" \
    --build-arg DOCKER_GO_BASE_TAG=latest \
    --build-arg DOCKER_ALPINE_BASE_TAG=latest \
    --build-arg STAGE_ENV="prod" .

echo "${SCRIPT_NAME} done."
