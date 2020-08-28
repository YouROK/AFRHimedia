#!/bin/bash

PLATFORMS=""
PLATFORMS_ARM="linux"

type setopt >/dev/null 2>&1

export GOPATH="${PWD}"

SCRIPT_NAME=`basename "$0"`
FAILURES=""
SOURCE_FILE="dist/afr"
CURRENT_DIRECTORY=${PWD##*/}
OUTPUT=${SOURCE_FILE:-$CURRENT_DIRECTORY} # if no src file given, use current dir name
LDFLAGS="-s -w"

GOARCH="arm"
GOARM="7"
GOOS="linux"

BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}${GOARM}"
CMD="GOARM=${GOARM} GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags=\"${LDFLAGS}\" -o ${OUTPUT} main"
echo "${CMD}"
eval "${CMD}" || FAILURES="${FAILURES} ${GOOS}/${GOARCH}${GOARM}" 

# eval errors
if [[ "${FAILURES}" != "" ]]; then
  echo ""
  echo "${SCRIPT_NAME} failed on: ${FAILURES}"
  exit 1
else
    adb push ./dist/afr /system/bin
fi
