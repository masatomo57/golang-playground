#!/bin/bash

go run main.go

ffplay -autoexit -f f32le -showmode 1 out.bin
