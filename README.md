# x-wallet

An experimental "web wallet client" playground for tfchain.

**WARNING**: This is an experimental playground, do not use it, unless you want to fool around. It does nothing useful at the moment.

## Experiments

### Pure Golang App

An experiment to show how to run a pure Golang app in the browser using WebAssembly.

Run it:

```
$ make rivine-wasm && caddy
GOARCH=wasm GOOS=js go1.11beta1 build -o app.wasm rivine/main.go
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

### Call Golang Callbacks from JS

An experiment to show how to define callbacks in Golang that can be called from JS.

> NOTE: this is not too useful for our purposes, as what we really want is to be able to define (**typed**) functions which can return results.

Run it:

```
$ make tfchain-wasm && caddy
GOARCH=wasm GOOS=js go1.11beta1 build -o app.wasm tfchain/main.go
Activating privacy features... done.
http://localhost:2015
```

Use your browser to surf to <http://localhost:2015>, open the browser console
and play with it (using the exposed JS global functions `randomkeypair` and `signtx`):

```
> randomkeypair()
< undefined
[Log] sk: 61a9e18e96f68fdbde65bc2754531ea5023925fb445e9158162657f4973874031c5bad93502a8106a84281b66eebb573bb556229d92a489ef29206ff347d9768 (go-wasm-runtime.js, line 257)
[Log] pk: 1c5bad93502a8106a84281b66eebb573bb556229d92a489ef29206ff347d9768 (go-wasm-runtime.js, line 257)
> signtx()
< undefined
[Log] signtx requires 3 params: an unsigned json tx, coin input index and hex-encoded private key (go-wasm-runtime.js, line 257)
> signtx('{"version":1,"data":{"coininputs":[{"parentid":"eac892a5a715b697ccafa6dd72a21fb7a3c1b3fad197890cc2b38a6ebf1bb9d3","fulfillment":{"type":1,"data":{"publickey":"ed25519:a474065da9eb7cfd75b6469e76251ab394dbb80f3312c5ee11370d660e3746fa"}}}],"coinoutputs":[{"value":"100000000000","condition":{}}],"minerfees":["100000000"]}}', 0, '3cc691b48219637c926a44eadb86fc8e8fd25fe5debbadb18ba6ca85247d4fb5a474065da9eb7cfd75b6469e76251ab394dbb80f3312c5ee11370d660e3746fa')
< undefined
[Log] {"version":1,"data":{"coininputs":[{"parentid":"eac892a5a715b697ccafa6dd72a21fb7a3c1b3fad197890cc2b38a6ebf1bb9d3","fulfillment":{"type":1,"data":{"publickey":"ed25519:a474065da9eb7cfd75b6469e76251ab394dbb80f3312c5ee11370d660e3746fa","signature":"8806cce8f137f352a99cdc9ab18ce8526eac8c5160ec4dd9f640e8af3a42ce05cbcddadcb3479730dfc5feb33aa251f76c872ff132c2887c73b58ff43a7d4105"}}}],"coinoutputs":[{"value":"100000000000","condition":{}}],"minerfees":["100000000"]}} (go-wasm-runtime.js, line 257)
```

## Repository Owners

* Rob Van Mieghem (@robvanmighem)
* Lee Smet (@leesmet)
* Glen De Cauwsemaecker (@glendc)
