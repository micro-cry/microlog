#!/bin/bash

scripts/git.sh --add_commit
scripts/git.sh --add_push

go mod tidy
#go generate ./...