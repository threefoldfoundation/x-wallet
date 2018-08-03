package main

import (
	"strconv"

	"github.com/rivine/rivine/crypto"
	"github.com/rivine/rivine/modules"
	"github.com/rivine/rivine/types"
)

func generateKeys(seed modules.Seed, start, n uint64) []spendableKey {
	// generate in parallel, one goroutine per core.
	keys := make([]spendableKey, n)
	for i := uint64(0); i < n; i++ {
		logPrint("generate spendable key for index " + strconv.FormatUint(start+i, 10))
		keys[i] = generateSpendableKey(seed, start+i)
	}
	return keys
}
func generateSpendableKey(seed modules.Seed, index uint64) spendableKey {
	sk, pk := crypto.GenerateKeyPairDeterministic(crypto.HashAll(seed, index))
	return spendableKey{
		PublicKey: pk,
		SecretKey: sk,
	}
}

type spendableKey struct {
	PublicKey crypto.PublicKey
	SecretKey crypto.SecretKey
}

func (sk spendableKey) UnlockHash() types.UnlockHash {
	return types.NewEd25519PubKeyUnlockHash(sk.PublicKey)
}
