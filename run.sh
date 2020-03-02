#!/usr/bin/env bash

ARCH=$1

if [ -z "${ARCH}" ]; then
  echo "need to specify arch"
  exit 1
fi

SERVER=term_${ARCH}

if [ "${ARCH}" == "win" ]; then
  SERVER=${SERVER}.exe
fi


## run the peripherals
printf "running the peripherals ... \n"

printf "display ... "
python3 display.py &
PID=$!
printf "[%d] done\n" "$PID"

printf "display ... "
python3 buzzer.py &
PID=$!
printf "[%d] done\n" "$PID"

printf "display ... "
python3 cardreader.py &
PID=$!
printf "[%d] done\n" "$PID"

printf "display ... "
python3 keyboard.py &
PID=$!
printf "[%d] done\n" "$PID"

## run the server
printf "running the server ... "
# shellcheck disable=SC2091
$( ./"${SERVER}" )
