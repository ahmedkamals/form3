ARG GO_VERSION
ARG ALPINE_VERSION
ARG MAINTAINER_NAME="Ahmed Kamal"
ARG MAINTAINER_EMAIL

FROM golang:${GO_VERSION:-1}-alpine${ALPINE_VERSION:-3.14} as builder
MAINTAINER ${MAINTAINER_NAME} <${MAINTAINER_EMAIL}>

WORKDIR /app

ARG API_ENDPOINT
ARG FIXTURES_PATH

ENV API_ENDPOINT=${API_ENDPOINT}
ENV FIXTURES_PATH=${FIXTURES_PATH}

COPY . .

RUN set -eux              \
    ;                     \
    apk update            \
    ;                     \
    apk upgrade           \
    ;                     \
    apk add --no-cache    \
        bash              \
        curl              \
        git               \
        make              \
    ;                     \
    make mods-download    \
    ; \
