#!/usr/bin/env bash

set -x
set -e

TARGET_IP=192.168.1.106

TARGET_OUT=/home/pi/

scp display.py pi@${TARGET_IP}:${TARGET_OUT}
