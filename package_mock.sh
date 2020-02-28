#!/usr/bin/env bash

set -e
#set -x

SWD=$( cd "$(dirname "${0}")" && pwd )
PERIPHS_SUBDIR=mock sh "${SWD}"/package.sh
