FROM alpine:3.17.0 AS deployment-container
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/btc-wallet"
LABEL stage="release"

# Configure least privilege user
ARG UID=1000
ARG GID=1000
RUN addgroup -S btcwallet -g ${UID} && \
    adduser -S btcwallet -u ${GID} -G btcwallet -h /home/btcwallet -s /bin/sh -D btcwallet

WORKDIR /

COPY btc-wallet /
COPY build/docker/btc-wallet/entrypoint.sh /

RUN mkdir -p /data/input && \
    mkdir -p /data/result && \
    mkdir -p /data/archive && \
    chown -R btcwallet:btcwallet /data

ENTRYPOINT ["/entrypoint.sh"]

USER btcwallet