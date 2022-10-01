package main

import (
	"golang-wasm/config"
	"golang-wasm/helpers"
	"golang-wasm/machine"
	"syscall/js"
)

func downloadDefaultEnigmaConfig() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		conf, _ := config.GenerateRandomEnigmaConfig(3)
		println(conf)
		buf ,_ := helpers.ConvertToBytesBuffer(conf)
		return helpers.GetJSReadableStreamFromIOReader(buf)
	})
}

func registerCallbacks() {
	js.Global().Set("downloadDefaultEnigmaConfig", downloadDefaultEnigmaConfig())
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
