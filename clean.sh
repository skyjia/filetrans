#!/bin/bash
set -e
set -u
set -o pipefail

CUR=$(dirname "$0")

TEMPLATE_DIR=${CUR}/testdata/template
TEST_DIR=${CUR}/testdata/test

if [ -d "${TEST_DIR}" ] 
then
    rm -rf "${TEST_DIR}"
fi

cp -r "${TEMPLATE_DIR}" "${TEST_DIR}"
