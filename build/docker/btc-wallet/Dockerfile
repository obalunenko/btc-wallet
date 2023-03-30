ARG DOCKER_REPO
ARG DOCKER_GO_BASE_TAG=latest
ARG DOCKER_ALPINE_BASE_TAG=latest
ARG STAGE_ENV="prod"
# hadolint ignore=DL3007
FROM ${DOCKER_REPO}btcwallet-go-base:${DOCKER_GO_BASE_TAG} AS build-container
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/btc-wallet"
LABEL stage=${STAGE_ENV}

ENV PROJECT_DIR="${GOPATH}/src/github.com/obalunenko/btc-wallet"

RUN mkdir -p "${PROJECT_DIR}"

WORKDIR "${PROJECT_DIR}"

COPY .git .git
COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum
COPY scripts scripts
COPY Makefile Makefile

# compile executable
RUN make build && \
    mkdir -p /app && \
    cp ./bin/btc-wallet /app/btc-wallet

COPY ./build/docker/btc-wallet/entrypoint.sh /app/entrypoint.sh

FROM ${DOCKER_REPO}btcwallet-alpine-base:${DOCKER_ALPINE_BASE_TAG} AS deployment-container
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/btc-wallet"
LABEL stage=${STAGE_ENV}

# Configure least privilege user
ARG UID=1000
ARG GID=1000
RUN addgroup -S btcwallet -g ${UID} && \
    adduser -S btcwallet -u ${GID} -G btcwallet -h /home/btcwallet -s /bin/sh -D btcwallet

WORKDIR /

COPY --from=build-container /app/ /

RUN mkdir -p /data/input && \
    mkdir -p /data/result && \
    mkdir -p /data/archive && \
    chown -R btcwallet:btcwallet /data

ENTRYPOINT ["sh", "-c", "/wait && /entrypoint.sh"]

USER btcwallet
