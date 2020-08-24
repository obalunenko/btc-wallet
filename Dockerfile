FROM golang:1.14-alpine as build-container

ENV PROJECT_DIR=${GOPATH}/src/github.com/oleg-balunenko/btc-wallet

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git musl-dev make gcc

RUN mkdir -p ${PROJECT_DIR}

COPY ./  ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}
# check vendor
RUN make gomod
# vet project
RUN make vet
# compile executable
RUN make compile

RUN mkdir /app
RUN cp ./bin/btc-wallet /app/btc-wallet


FROM alpine:3.11.3 as deployment-container
RUN apk add -U --no-cache ca-certificates


COPY --from=build-container /app/btc-wallet /btc-wallet

ENTRYPOINT ["/btc-wallet"]

