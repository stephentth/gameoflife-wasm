.PHONY: build, test

build:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" docs/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o docs/script.wasm wasm/*.go

test:
	go test ./algorithm
