#! /usr/bin/env bash

cd webserver
GOOS=js GOARCH=wasm go build -o static/script.wasm wasm/*.go
go build -o server
./server