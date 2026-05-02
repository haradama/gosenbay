//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"strconv"
	"syscall/js"

	"github.com/haradama/gosenbay/senbay"
)

const defaultPN = 121

// encode is a function that encodes a JSON string into a Senbay string.
func encode(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return ""
	}

	compress := true
	if len(args) >= 2 {
		compress = args[1].Bool()
	}

	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(args[0].String()), &payload); err != nil {
		return ""
	}

	sd, err := senbay.NewSenbayData(defaultPN)
	if err != nil {
		return ""
	}

	for key, value := range payload {
		switch v := value.(type) {
		case float64:
			sd.AddFloat64(key, v)
		case string:
			sd.AddText(key, v)
		case bool:
			sd.AddText(key, strconv.FormatBool(v))
		}
	}

	return sd.Encode(compress)
}

// decode is a function that decodes a Senbay string into a JSON string.
func decode(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return "{}"
	}

	sd, err := senbay.NewSenbayData(defaultPN)
	if err != nil {
		return "{}"
	}

	decoded := sd.Decode(args[0].String())

	bytes, err := json.Marshal(decoded)
	if err != nil {
		return "{}"
	}

	return string(bytes)
}

func main() {
	js.Global().Set("senbayEncode", js.FuncOf(encode))
	js.Global().Set("senbayDecode", js.FuncOf(decode))

	select {}
}
