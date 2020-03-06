#!/usr/bin/env bash

SWD=$( cd "$(dirname "${0}")" && pwd )

scp -r "${SWD}/out" pi@192.168.1.101:/home/pi/payterm
