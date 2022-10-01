package main

import (
	"bytes"
	"encoding/json"
	"golang-wasm/helpers"
	"golang-wasm/machine"
	"golang-wasm/rotor"
	"syscall/js"
)

func downloadDefaultConfig() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := new(bytes.Buffer)
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(rotor.DefaultConfig)
		println("err", err)
		return helpers.GetJSReadableStreamFromIOReader(b)
	})
}

func registerCallbacks() {
	js.Global().Set("downloadDefaultConfig", downloadDefaultConfig())
	js.Global().Set("ScrambleCharacter", js.FuncOf(func(this js.Value, args []js.Value) any {

		m := machine.NewMachine(nil)

		str := ""
		for _, v := range "[Q Q;1P $= _; *\"-)1, |8OMA5>Y; %}[ALU)6" {
			str += m.ScrambleCharacter(string(v))
		}

		return str
	}))
}

func main() {
	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	select {}
}
