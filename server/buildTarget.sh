#!/usr/bin/env bash

set -x
set -e

TARGET_IP=192.168.1.106
TARGET_OS=linux
TARGET_ARCH=arm

OUT=out/payterm
TARGET_OUT=/home/pi/payterm

GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -o ${OUT} cmd/main.go

scp ${OUT} pi@${TARGET_IP}:${TARGET_OUT}
# shellcheck disable=SC2029
ssh pi@${TARGET_IP} "chmod +x ${TARGET_OUT}"
