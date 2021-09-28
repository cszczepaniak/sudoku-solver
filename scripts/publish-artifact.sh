set -ex

aws s3 cp $GITHUB_SHA.zip s3://sudoku-solver-artifacts/