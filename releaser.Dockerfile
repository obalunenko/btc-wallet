FROM alpine:3.17.0
RUN apk add -U --no-cache ca-certificates

COPY btc-wallet /

ENTRYPOINT ["/btc-wallet"]