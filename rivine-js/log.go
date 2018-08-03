package main

import "github.com/gopherjs/gopherjs/js"

func logPrint(str string) {
	js.Global.Get("console").Call("log", str)
}
