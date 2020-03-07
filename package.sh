#!/usr/bin/env bash

set -e
#set -x

SWD=$( cd "$(dirname "${0}")" && pwd )
OUT_DIR=${SWD}/out
CONTROLLER_DIR=${SWD}/controller
PERIPHS_DIR=${SWD}/peripherals


rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

pushd "${CONTROLLER_DIR}" > /dev/null
printf "refreshing go modules ... "
go mod download
printf "done\n"

printf "building terminal controller x64 linux ... "
GOOS=linux GOARCH=amd64 go build -o "${OUT_DIR}"/controller_x64 cmd/controller.go
printf "done\n"

printf "building terminal controller arm linux ... "
GOOS=linux GOARCH=arm go build -o "${OUT_DIR}"/controller_arm cmd/controller.go
printf "done\n"

printf "building terminal controller x64 windows ... "
GOOS=windows GOARCH=amd64 go build -o "${OUT_DIR}"/controller_win.exe cmd/controller.go
printf "done\n"
popd > /dev/null

printf "copying peripheral controllers ... "
rsync -a --exclude '__pycache__' "${PERIPHS_DIR}"/. "${OUT_DIR}"
printf "done\n"

printf "copying runner script ... "
cp "${SWD}"/run.sh "${OUT_DIR}"
printf "done\n"

printf "\n"

ls -l "${OUT_DIR}"
