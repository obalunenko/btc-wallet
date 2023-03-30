FROM alpine:3.17.0
LABEL maintainer="oleg.balunenko@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/obalunenko/btc-wallet"

LABEL stage="base"

ARG APK_CA_CERTIFICATES_VERSION=20220614-r4
RUN apk update && \
    apk add --no-cache \
        "ca-certificates=${APK_CA_CERTIFICATES_VERSION}" && \
    rm -rf /var/cache/apk/*

WORKDIR /

## Add the wait script to the image
ARG WAIT_VERSION=2.10.0
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/${WAIT_VERSION}/wait /wait
RUN chmod +x /wait
