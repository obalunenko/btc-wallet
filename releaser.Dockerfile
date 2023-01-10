FROM alpine:3.17.1
RUN apk add -U --no-cache ca-certificates

COPY btc-wallet /

ENTRYPOINT ["/btc-wallet"]