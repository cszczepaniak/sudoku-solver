# Sudoku Solver
This is a simple sudoku solver. My goal is to also build two frontends for the core logic: a CLI and a REST API.

## Running
Run `make cli` to start to CLI, or `make rest` to start the REST server locally.

## Building
There is an automated pipeline that will build, zip, and upload the lambda to S3. It will then update the lambda in AWS. See `scripts/` for details.

## What's in `/rust`?
I originally built this in Rust, and my original solution is there. However, I'm not an experienced Rust programmer (yet), so it was easier for me to switch to Go to make something I could deploy more easily and quickly.
