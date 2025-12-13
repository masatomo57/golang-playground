#!/bin/bash
set -e
title=${1:-jingle_bell}

go run main.go -title=$title

ffplay -autoexit -f f32le -showmode 1 out.bin
