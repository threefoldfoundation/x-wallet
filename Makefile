rivine-wasm:
	GOARCH=wasm GOOS=js go1.11beta1 build -o app.wasm rivine/main.go

tfchain-wasm:
	GOARCH=wasm GOOS=js go1.11beta1 build -o app.wasm tfchain/main.go