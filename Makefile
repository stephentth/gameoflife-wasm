.PHONY: build, test

build:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" web/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o web/script.wasm wasm/*.go

test:
	go test ./algorithm
