#!/bin/bash

docker run -it \
  -v ./:/go/src/hello-wasi \
  golang:1.21-bookworm \
  bash

