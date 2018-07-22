#!/usr/bin/env bash
set -e

go build
./music504 --enable_tips
