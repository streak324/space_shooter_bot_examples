package main

import "unsafe"

type Entity struct {
	My    bool   `json:"my"`
	Id    uint64 `json:"id"`
	Owner int32  `json:"owner"`
}

type GameState struct {
	Entities []Entity `json:"entities"`
}

var printBuffer [1024]byte

//go:wasmimport host printInt
func printInt(integer int32)

//go:wasmimport host print
func print(ptr uint32, len uint32)

//go:export step
func step() {
	printInt(1235)
	copyLength := copy(printBuffer[:], []byte("Hello WASM Host!!!"))
	intptr := uint32(uintptr(unsafe.Pointer(&printBuffer[0])))
	print(intptr, uint32(copyLength))
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
