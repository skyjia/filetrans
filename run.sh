#!/bin/bash
set -e
set -u
set -o pipefail

CUR=$(dirname "$0")
TEST_DIR=${CUR}/testdata/test

go run main.go \
    -app_id=20210707000882113 \
    -key=_xfaR1zsBEOk7H1Yi60b \
    -delay=1000 \
    -p=${TEST_DIR} \
    -t=.txt
