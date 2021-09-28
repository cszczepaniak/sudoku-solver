#!/bin/bash
set -ex

source ./scripts/variables.sh

aws s3 cp $BUILD_NAME.zip s3://$ARTIFACT_BUCKET/