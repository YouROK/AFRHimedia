#!/bin/bash

./build.sh && adb push dist/afr /data/local/tmp && adb shell "killall afr;/data/local/tmp/afr"
