#!/bin/bash
set -ex

source ./scripts/variables.sh

aws update-function-code --function-name=$LAMBDA_NAME --s3-bucket=$ARTIFACT_BUCKET --s3-key=$BUILD_NAME.zip