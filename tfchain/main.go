package main

import (
	"encoding/hex"
	"encoding/json"
	"strconv"
	"syscall/js"

	"github.com/rivine/rivine/crypto"
	"github.com/rivine/rivine/encoding"
	"github.com/rivine/rivine/types"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("decodetx", js.NewCallback(decodeTx))
	js.Global().Set("randomkeypair", js.NewCallback(randomKeyPair))
	js.Global().Set("signtx", js.NewCallback(signTx))
	<-c
}

func randomKeyPair([]js.Value) {
	sk, pk := crypto.GenerateKeyPair()
	js.Global().Get("console").Call("log", "sk: "+hex.EncodeToString(sk[:]))
	js.Global().Get("console").Call("log", "pk: "+hex.EncodeToString(pk[:]))
}

func decodeTx(vars []js.Value) {
	if len(vars) == 0 {
		js.Global().Get("console").Call("log", "decodetx requires 1 param: a hex-encoded transaction")
		return
	}
	b, err := hex.DecodeString(vars[0].String())
	if err != nil {
		js.Global().Get("console").Call("log", "decodetx failed to hex-decode string to byteslice: "+err.Error())
		return
	}
	var tx types.Transaction
	err = encoding.Unmarshal(b, &tx)
	if err != nil {
		js.Global().Get("console").Call("log", "decodetx failed to unmarshal tx: "+err.Error())
		return
	}
	b, _ = json.Marshal(tx)
	js.Global().Get("console").Call("log", string(b))
}

func signTx(vars []js.Value) {
	if len(vars) != 3 {
		js.Global().Get("console").Call("log", "signtx requires 3 params: an unsigned json tx, coin input index and hex-encoded private key")
		return
	}

	var tx types.Transaction
	err := json.Unmarshal([]byte(vars[0].String()), &tx)
	if err != nil {
		js.Global().Get("console").Call("log", "signTx: failed to json-unmarshal json tx: "+err.Error())
		return
	}

	b, err := hex.DecodeString(vars[2].String())
	if err != nil {
		js.Global().Get("console").Call("log", "signTx: failed to hex-decode private key to byteslice: "+err.Error())
		return
	}
	if len(b) != crypto.SecretKeySize {
		js.Global().Get("console").Call("log", "signTx: private key has wrong size: "+strconv.Itoa(len(b)))
		return
	}
	var sk crypto.SecretKey
	copy(sk[:], b[:])

	index := vars[1].Int()
	err = tx.CoinInputs[index].Fulfillment.Sign(types.FulfillmentSignContext{
		InputIndex:  uint64(index),
		Transaction: tx,
		Key:         sk,
	})
	if err != nil {
		js.Global().Get("console").Call("log", "signTx: error while signing: "+err.Error())
		return
	}

	b, _ = json.Marshal(tx)
	js.Global().Get("console").Call("log", string(b))
}
