FROM golang:1.16.5-alpine as build-container

ENV PROJECT_DIR=${GOPATH}/src/github.com/obalunenko/btc-wallet

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
RUN cp ./scripts/entrypoint.sh  /app/entrypoint.sh


FROM alpine:3.15.3 as deployment-container
RUN apk add -U --no-cache ca-certificates

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

RUN mkdir -p logs


COPY --from=build-container /app/btc-wallet /btc-wallet
COPY --from=build-container /app/entrypoint.sh /entrypoint.sh

ENTRYPOINT /wait && /entrypoint.sh

