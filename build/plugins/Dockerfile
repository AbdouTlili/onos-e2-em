# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0 as pluginbuild
# FROM ubuntu:latest

ENV GO111MODULE=on
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION

COPY build/_input/${PLUGIN_MAKE_TARGET} /build

WORKDIR /build


RUN GOFLAGS=-mod=vendor CGO_ENABLED=1 go build -o ${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} \
    -buildmode=plugin .

FROM alpine:3.12
ARG PLUGIN_MAKE_TARGET
ARG PLUGIN_MAKE_VERSION
COPY --from=pluginbuild /build/${PLUGIN_MAKE_TARGET}.so.${PLUGIN_MAKE_VERSION} /

