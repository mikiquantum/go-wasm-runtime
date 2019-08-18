package main

import (
	"log"

	wasmgo "github.com/mikiquantum/go-wasm-runtime"
)

func main() {
	b, err := wasmgo.BridgeFromFile("test", "./simple/prog/main.wasm", nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := b.Run(); err != nil {
		log.Fatal(err)
	}
}
