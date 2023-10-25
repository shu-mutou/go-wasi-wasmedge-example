#!/bin/bash

docker run -it \
  -v ./:/go/src/hello-wasi \
  golang:1.21-bookworm \
  GOOS=wasip1 GOARCH=wasm go build -o main.wasi main.go

