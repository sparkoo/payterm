#!/usr/bin/env bash

set -e
#set -x

SWD=$( cd "$(dirname "${0}")" && pwd )
OUT_DIR=${SWD}/out
TERM_DIR=${SWD}/term
PERIPHS_DIR=${SWD}/peripherals

if [ -z "${PERIPHS_SUBDIR}" ]; then
  PERIPHS_SUBDIR="."
fi

rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

pushd "${TERM_DIR}" > /dev/null
printf "building terminal server ... "
go mod download
go build -v -o "${OUT_DIR}"/term cmd/term.go
printf "done\n"
popd > /dev/null

pushd "${PERIPHS_DIR}" > /dev/null
printf "copying peripheral controllers ... "
cp "${PERIPHS_SUBDIR}"/*.py "${OUT_DIR}"
cp -r lib/ "${OUT_DIR}"
printf "done\n"
popd > /dev/null

printf "\n"

ls -l "${OUT_DIR}"
