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
DISPLAY_PID=$!
printf "[%d] done\n" "$DISPLAY_PID"

printf "display ... "
python3 buzzer.py &
BUZZER_PID=$!
printf "[%d] done\n" "$BUZZER_PID"

printf "display ... "
python3 cardreader.py &
CARDREADER_PID=$!
printf "[%d] done\n" "$CARDREADER_PID"

printf "display ... "
python3 keyboard.py &
KEYBOARD_PID=$!
printf "[%d] done\n" "$KEYBOARD_PID"

## run the server
printf "running the server ... "
# shellcheck disable=SC2091
$( ./"${SERVER}" )
