# x-wallet

An experimental web wallet client for tfchain.

**WARNING**: This is an experimental project, still in development.

## how to run locally

```
$ make rivine-wasm && caddy
GOARCH=wasm GOOS=js go1.11beta1 build -o rivine.wasm rivine/main.go
Activating privacy features... done.
http://localhost:2015
```

Use your browser to surf to <http://localhost:2015>, open the browser console and you'll find the following computed results there:

```
[Log] Unsigned Transaction in HEX AND JSON: (go-wasm-runtime.js, line 257)
[Log] {"version":1,"data":{"coininputs":[{"parentid":"eac892a5a715b697ccafa6dd72a21fb7a3c1b3fad197890cc2b38a6ebf1bb9d3","fulfillment":{"type":1,"data":{"publickey":"ed25519:6a1a850c9d4cd28b9adca3641c6dec9766c8a0f0ae4a77efb43a291d5f9e510f","signature":""}}}],"coinoutputs":[{"value":"100000000000","condition":{}}],"minerfees":["100000000"]}} (go-wasm-runtime.js, line 257)
[Log] 01bb000000000000000100000000000000eac892a5a715b697ccafa6dd72a21fb7a3c1b3fad197890cc2b38a6ebf1bb9d30140000000000000006564323535313900000000000000000020000000000000006a1a850c9d4cd28b9adca3641c6dec9766c8a0f0ae4a77efb43a291d5f9e510f000000000000000001000000000000000500000000000000174876e800000000000000000000000000000000000000000000000000000100000000000000040000000000000005f5e1000000000000000000 (go-wasm-runtime.js, line 257)
[Log] Signed Transaction in HEX AND JSON: (go-wasm-runtime.js, line 257)
[Log] {"version":1,"data":{"coininputs":[{"parentid":"eac892a5a715b697ccafa6dd72a21fb7a3c1b3fad197890cc2b38a6ebf1bb9d3","fulfillment":{"type":1,"data":{"publickey":"ed25519:6a1a850c9d4cd28b9adca3641c6dec9766c8a0f0ae4a77efb43a291d5f9e510f","signature":"31dc2cade36e04aae24aa5c281ba0383dbc1da65a53f8cdd6165368222253359982b7540275e3b6601d426505ca68bf26bd9103c3365426c937209ad9456c604"}}}],"coinoutputs":[{"value":"100000000000","condition":{}}],"minerfees":["100000000"]}} (go-wasm-runtime.js, line 257)
[Log] 01fb000000000000000100000000000000eac892a5a715b697ccafa6dd72a21fb7a3c1b3fad197890cc2b38a6ebf1bb9d30180000000000000006564323535313900000000000000000020000000000000006a1a850c9d4cd28b9adca3641c6dec9766c8a0f0ae4a77efb43a291d5f9e510f400000000000000031dc2cade36e04aae24aa5c281ba0383dbc1da65a53f8cdd6165368222253359982b7540275e3b6601d426505ca68bf26bd9103c3365426c937209ad9456c60401000000000000000500000000000000174876e800000000000000000000000000000000000000000000000000000100000000000000040000000000000005f5e1000000000000000000 (go-wasm-runtime.js, line 257)
```
