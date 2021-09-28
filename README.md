# Sudoku Solver
This is a simple sudoku solver. My goal is to also build two frontends for the core logic: a CLI and a REST API.

## Running
Run `make cli` to start to CLI, or `make rest` to start the REST server locally.

## Building
There is an automated pipeline that will build, zip, and upload the lambda to S3. It will then update the lambda in AWS. See `scripts/` for details.

## What's in `/rust`?
I originally built this in Rust, and my original solution is there. I keep going back to Rust to learn a little more about it. I rebuilt it in Go because deploying a REST API to AWS Lambda/API Gateway is _so simple_ in Go but seemingly not quite as simple in Rust. I did have fun with Rust `enum`s and `match` for this, but ultimately Go was a better choice to get up and running in the cloud for me.
