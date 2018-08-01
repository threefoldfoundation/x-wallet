package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/rivine/rivine/crypto"
	"github.com/rivine/rivine/encoding"
	"github.com/rivine/rivine/types"
)

func main() {
	// create random key pair
	// test public/private key pair
	sk, pk := crypto.GenerateKeyPair()
	ed25519pk := types.Ed25519PublicKey(pk)

	var randomParentID types.CoinOutputID
	rand.Read(randomParentID[:])

	// hard-coded tx
	tx := types.Transaction{
		Version: types.TransactionVersionOne,
		CoinInputs: []types.CoinInput{
			{
				ParentID:    randomParentID,
				Fulfillment: types.NewFulfillment(&types.SingleSignatureFulfillment{PublicKey: ed25519pk}),
			},
		},
		CoinOutputs: []types.CoinOutput{
			{
				Value:     types.NewCurrency64(100000000000),
				Condition: types.NewCondition(&types.NilCondition{}),
			},
		},
		MinerFees: []types.Currency{types.NewCurrency64(100000000)},
	}

	// display unsigned tx
	js.Global().Get("console").Call("log", "Unsigned Transaction in HEX AND JSON:")
	js.Global().Get("console").Call("log", func() string { b, _ := json.Marshal(tx); return string(b) }())
	js.Global().Get("console").Call("log", fmt.Sprintf("%x", encoding.Marshal(tx)))

	// sign tx
	err := tx.CoinInputs[0].Fulfillment.Sign(types.FulfillmentSignContext{
		InputIndex:  0,
		Transaction: tx,
		Key:         sk,
	})
	if err != nil {
		js.Global().Get("console").Call("log", "Failed to sign: "+err.Error())
		return
	}

	// display signed tx
	js.Global().Get("console").Call("log", "Signed Transaction in HEX AND JSON:")
	js.Global().Get("console").Call("log", func() string { b, _ := json.Marshal(tx); return string(b) }())
	js.Global().Get("console").Call("log", fmt.Sprintf("%x", encoding.Marshal(tx)))
}
