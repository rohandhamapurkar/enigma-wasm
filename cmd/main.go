package main

import (
	"encoding/json"
	"golang-wasm/config"
	"golang-wasm/helpers"
	"golang-wasm/machine"
	"syscall/js"
)

func getDefaultEnigmaConfigStream() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		conf := config.DefaultEnigmaConfig
		buf, _ := helpers.ConvertToBytesBuffer(conf)
		return helpers.GetJSReadableStreamFromIOReader(buf)
	})
}

func getCurrentConfigStream(m *machine.Machine) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		buf, _ := m.EnigmaConfig.ToJSONBytesBuffer()
		return helpers.GetJSReadableStreamFromIOReader(buf)
	})
}

func setEnigmaConfig(m *machine.Machine) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		rc := &config.EnigmaConfig{}
		err := json.Unmarshal([]byte(args[0].String()), rc)
		if(err != nil) {
			println(err)
		}
		m.SetConfig(rc)
		return true
	})
}

func registerCallbacks() {
	m := machine.NewMachine(nil)
	js.Global().Set("getDefaultEnigmaConfigStream", getDefaultEnigmaConfigStream())
	js.Global().Set("getCurrentConfigStream", getCurrentConfigStream(m))
	js.Global().Set("setEnigmaConfig", setEnigmaConfig(m))
	js.Global().Set("ScrambleText", js.FuncOf(func(this js.Value, args []js.Value) any {

		str := ""
		for _, char := range args[0].String() {
			str += m.ScrambleCharacter(string(char))
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
