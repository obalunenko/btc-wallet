FROM alpine:3.14.2
RUN apk add -U --no-cache ca-certificates

COPY btc-wallet /

ENTRYPOINT ["/btc-wallet"]