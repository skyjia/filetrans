#!/bin/bash
set -e
set -u
set -o pipefail

CUR=$(dirname "$0")

TEMPLATE_DIR=${CUR}/testdata/template
TEST_DIR=${CUR}/testdata/test

rm -rf ${TEST_DIR}
cp -r ${TEMPLATE_DIR} ${TEST_DIR}
