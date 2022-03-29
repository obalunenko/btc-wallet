FROM alpine:3.15.3
RUN apk add -U --no-cache ca-certificates

COPY btc-wallet /

ENTRYPOINT ["/btc-wallet"]