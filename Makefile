rivine-wasm:
	GOARCH=wasm GOOS=js go1.11beta1 build -o rivine.wasm rivine/main.go