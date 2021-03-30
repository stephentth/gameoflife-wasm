.PHONY: build, test

build: generate
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" docs/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o docs/script.wasm wasm/*.go

generate:
	go run ./patterns_embed/generator.go 
	mv assets_vfsdata.go patterns_embed/

test:
	go test ./algorithm
