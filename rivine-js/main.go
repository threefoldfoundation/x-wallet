package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("rivine", map[string]interface{}{
		"account": map[string]interface{}{
			"New": NewAccount,
		},
		"address": map[string]interface{}{
			"New": NewAddress,
		},
		"wallet": map[string]interface{}{
			"New": NewWallet,
		},
	})
}
