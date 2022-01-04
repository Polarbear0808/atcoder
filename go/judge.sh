#!/bin/bash
set -eu
ANSWER=$1
DATASET_DIR=$2

for file in $DATASET_DIR/*.txt; do
    go run $ANSWER < $file
done