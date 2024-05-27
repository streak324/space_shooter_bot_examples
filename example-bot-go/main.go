package main

import (
	"bot-go/gjson"
	"runtime"
	"strconv"
	"unsafe"
)

// some code taken from https://github.com/tetratelabs/wazero/blob/1e0f88bc1462ca07a33df83004914d3af7f5bcb4/examples/allocation/tinygo/testdata/greet.go

var printBuffer [1024]byte

var jsonGameStateBuffer [4096]byte

func log(message string) {
	ptr, size := stringToPtr(message)
	_log(ptr, size)
	runtime.KeepAlive(message) // keep message alive until ptr is no longer needed.
}

// _log is a WebAssembly import which prints a string (linear memory offset,
// byteCount) to the console.
//
//go:wasmimport env log
func _log(ptr, size uint32)

// stringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
// The returned pointer aliases the string hence the string must be kept alive
// until ptr is no longer needed.
func stringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

func bytesToPtr(b []byte) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.SliceData(b))
	return uint32(uintptr(ptr)), uint32(len(b))
}

// ptrToString returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func ptrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

//go:export getJSONGameState
func getJSONGameState() []byte {
	size := _getJSONGameState(bytesToPtr(jsonGameStateBuffer[:cap(jsonGameStateBuffer)]))
	return jsonGameStateBuffer[:size]
}

//go:wasmimport env getJSONGameState
func _getJSONGameState(ptr uint32, capacity uint32) int32

//go:export step
func step() {
	log("Hello WASM Host!!!")
	log(strconv.Itoa(5678))
	//log(fmt.Sprintf("%d", 1337))

	jsonData := ptrToString(bytesToPtr(getJSONGameState()))
	entitiesJSON := gjson.Get(jsonData, "entities").Array()
	for _, entityJSON := range entitiesJSON {
		my := entityJSON.Get("my").Bool()
		id := entityJSON.Get("id").Int()
		if my {
			log("id: " + strconv.FormatInt(id, 10) + " is mine")
		} else {
			log("id: " + strconv.FormatInt(id, 10) + " is not mine")
		}
	}
	log(jsonData)

	//data, err := easyjson.Marshal(serialize.GameState{
	//	Entities: []serialize.Entity{
	//		{
	//			My:    true,
	//			Id:    1,
	//			Owner: 2,
	//		},
	//	}})
	//if err != nil {
	//	copyLength := copy(printBuffer[:], []byte(err.Error()))
	//	intptr := uint32(uintptr(unsafe.Pointer(&printBuffer[0])))
	//	print(intptr, uint32(copyLength))
	//	return
	//}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
