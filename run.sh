#!/usr/bin/env bash
set -e

go build
./music504 --time_period=$1
