#!/bin/bash
set -ex

BUILD_NAME=$GITHUB_SHA
ARTIFACT_BUCKET="sudoku-solver-artifacts"
LAMBDA_NAME="sudoku-solver"