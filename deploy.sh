#!/usr/bin/env bash

SWD=$( cd "$(dirname "${0}")" && pwd )

TARGET_IP=192.168.1.101

ssh pi@${TARGET_IP} "rm -rf /home/pi/payterm"
scp -r "${SWD}/out" pi@${TARGET_IP}:/home/pi/payterm
