#!/usr/bin/env bash

set -e
#set -x

SWD=$( cd "$(dirname "${0}")" && pwd )
OUT_DIR=${SWD}/out
TERM_DIR=${SWD}/term
PERIPHS_DIR=${SWD}/peripherals


rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

pushd "${TERM_DIR}" > /dev/null
printf "refreshing go modules ... "
go mod download
printf "done\n"

printf "building terminal server x64 linux ... "
GOOS=linux GOARCH=amd64 go build -o "${OUT_DIR}"/term_x64 cmd/term.go
printf "done\n"

printf "building terminal server arm linux ... "
GOOS=linux GOARCH=arm go build -o "${OUT_DIR}"/term_arm cmd/term.go
printf "done\n"

printf "building terminal server x64 windows ... "
GOOS=windows GOARCH=amd64 go build -o "${OUT_DIR}"/term_win.exe cmd/term.go
printf "done\n"
popd > /dev/null

printf "copying peripheral controllers ... "
rsync -a --exclude '__pycache__' "${PERIPHS_DIR}"/. "${OUT_DIR}"
printf "done\n"

printf "\n"

ls -l "${OUT_DIR}"
