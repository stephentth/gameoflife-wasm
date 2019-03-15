#! /usr/bin/env bash

if [[ ! -f "$(go env GOROOT)/misc/wasm/wasm_exec.js" ]]; then
    echo "Recheck your develope environment for compatible Go version"
    exit 1
fi

if ! diff webserver/static/wasm_exec.js "$(go env GOROOT)/misc/wasm/wasm_exec.js"; then
    echo "New wasm_exec.js detected. Copy new file."
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" webserver/static/wasm_exec.js
fi

cd webserver
GOOS=js GOARCH=wasm go build -o static/script.wasm wasm/*.go
go build -o server
./server