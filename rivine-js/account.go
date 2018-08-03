package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/rivine/rivine/api"
	"github.com/rivine/rivine/modules"
	"github.com/rivine/rivine/types"
)

type (
	Account struct {
		Seed            modules.Seed
		Wallet          Wallet
		Keys            map[types.UnlockHash]spendableKey
		MultisigWallets map[types.UnlockHash]*MultiSigWallet

		mux sync.RWMutex
	}
	Wallet struct {
		address        types.UnlockHash
		unspendOutputs []types.CoinOutput
	}
	MultiSigWallet struct {
		Wallet
		UnlockHashes          []types.UnlockHash
		MinimumSignatureCount uint64
	}
)

func NewAccount(mnemonic string) *js.Object {
	seed, err := modules.InitialSeedFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}
	return js.MakeWrapper(&Account{
		Seed:            seed,
		Keys:            make(map[types.UnlockHash]spendableKey),
		MultisigWallets: make(map[types.UnlockHash]*MultiSigWallet),
	})
}

// Scan all info for this account
func (account *Account) Scan(explorerAddress string) {
	// TODO: broken and slow
	if explorerAddress == "undefined" {
		panic("no explorer address given")
	}

	account.mux.Lock()
	defer account.mux.Unlock()

	const (
		initialKeys = 2025 // based on the original key size limitations of rivine
		keyHop      = 1024
		maxKeys     = 100e6
	)

	type (
		unlockHashIndexPair struct {
			UnlockHash types.UnlockHash
			Index      uint64
		}
		scanResult struct {
			MultiSigAddresses []types.UnlockHash
			UnlockHashFound   bool
			KeyIndex          uint64
			ChannelCloseIndex uint64
		}
	)

	keys := generateKeys(account.Seed, 0, initialKeys)
	highestKeyIndex := uint64(0)
	ki := uint64(0)
	foundKeys := false

	// scan keys and gather results
	for {
		// check our key index, and expand key range if needed
		if length := uint64(len(keys)); ki == length {
			// check foundKeys & ki
			if !foundKeys {
				// no keys were found, we can stop
				break
			}
			if ki == maxKeys {
				// stop the scanning with a panic
				panic("reached maximum amount of keys, and still scanning")
			}
			// generate more keys
			newLength := length + keyHop
			if newLength > maxKeys {
				// clamp to max keys
				newLength = maxKeys
			}
			keys = append(keys, generateKeys(account.Seed, ki, newLength-length)...)
			foundKeys = false
		}

		uh := keys[ki].UnlockHash()

		resp, err := http.Get(explorerAddress + "/explorer/hashes/" + uh.String())
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			if resp.StatusCode != 400 {
				msg, _ := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				panic(string(msg))
			}
			logPrint("ignorning uh (" + strconv.FormatUint(ki, 10) + ") " +
				uh.String() + " as we received status code " + strconv.Itoa(resp.StatusCode))
			continue
		}
		decoder := json.NewDecoder(resp.Body)
		var explorerHashGet api.ExplorerHashGET
		err = decoder.Decode(&explorerHashGet)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}

		highestKeyIndex++

		mn := len(explorerHashGet.MultiSigAddresses)
		if mn == 0 {
			logPrint("uh (" + strconv.FormatUint(ki, 10) + ") " +
				uh.String() + " is not part of any multisig wallet")
			continue
		}
		logPrint("uh (" + strconv.FormatUint(ki, 10) + ") " +
			uh.String() + " is part of " + strconv.Itoa(mn) + " multisig wallet(s)")
		// TODO: store key address
	}

	for _, key := range keys[:highestKeyIndex] {
		account.Keys[key.UnlockHash()] = key
	}
}

// Addresses returns all the addresses of the individual wallet of this account
func (account *Account) Addresses() (addresses []string) {
	account.mux.RLock()
	defer account.mux.RUnlock()
	for _, key := range account.Keys {
		addresses = append(addresses, key.UnlockHash().String())
	}
	return
}

// NewAddress generates a new address for a given mnemonic and index
func NewAddress(mnemonic string, index uint64) string {
	seed, err := modules.InitialSeedFromMnemonic(mnemonic)
	if err != nil {
		panic(err)
	}
	return generateSpendableKey(seed, index).UnlockHash().String()
}

func NewWallet(address string) *js.Object {
	var uh types.UnlockHash
	err := uh.LoadString(address)
	if err != nil {
		panic(err)
	}
	return js.MakeWrapper(&Wallet{
		address: uh,
	})
}

func (w *Wallet) Address() string {
	return w.address.String()
}

func (w *Wallet) MultiSigAddresses() {
	// TODO: no blocking calls are allowed, what to do with funcs such as this?
	go func() {
		resp, err := http.Get("https://explorer.testnet.threefoldtoken.com/explorer/hashes/" + w.Address())
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			if resp.StatusCode != 400 {
				msg, _ := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				panic(string(msg))
			}
		}
		decoder := json.NewDecoder(resp.Body)
		var explorerHashGet api.ExplorerHashGET
		err = decoder.Decode(&explorerHashGet)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}
		var addresses []string
		for _, uh := range explorerHashGet.MultiSigAddresses {
			addresses = append(addresses, uh.String())
		}
		logPrint(strings.Join(addresses, ", "))
	}()
	return
}

// Address returns the MultiSig address of this MultiSigWallet
func (msw *MultiSigWallet) Address() string {
	return (&types.MultiSignatureCondition{
		UnlockHashes:          msw.UnlockHashes,
		MinimumSignatureCount: msw.MinimumSignatureCount,
	}).UnlockHash().String()
}
