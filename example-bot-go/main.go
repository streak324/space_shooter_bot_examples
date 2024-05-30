package main

import (
	"bot-go/gjson"
	"bot-go/serialize"
	"math"
	"runtime"
	"strconv"
	"unsafe"
)

// some code taken from https://github.com/tetratelabs/wazero/blob/1e0f88bc1462ca07a33df83004914d3af7f5bcb4/examples/allocation/tinygo/testdata/greet.go

var printBuffer []byte = make([]byte, 2048)

var jsonGameStateBuffer [4096]byte

func log(message string) {
	ptr, size := stringToPtr(message)
	_log(ptr, size)
	runtime.KeepAlive(message) // keep message alive until ptr is no longer needed.
}

func logb(message []byte) {
	ptr, size := bytesToPtr(message)
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

//go:wasmimport env moveEntityToTarget
func moveEntityToTarget(entityId uint64, x float32, y float32) int32

var stepCount uint64 = 0
var gotoPoints [2]serialize.Vec2
var gotoIndex = 0

//go:export step
func step() {
	jsonData := ptrToString(bytesToPtr(getJSONGameState()))
	entitiesJSON := gjson.Get(jsonData, "entities").Array()
	for _, entityJSON := range entitiesJSON {
		my := entityJSON.Get("my").Bool()
		id := entityJSON.Get("id").Int()
		posX := float32(entityJSON.Get("position").Get("x").Float())
		posY := float32(entityJSON.Get("position").Get("y").Float())
		if my {
			wholeX, fractionX := math.Modf(float64(posX))
			integerWholeX := int(math.Floor(wholeX))
			integerFractionX := int(math.Abs(math.Floor(fractionX * 1000)))

			wholeY, fractionY := math.Modf(float64(posY))
			integerWholeY := int(math.Floor(float64(wholeY)))
			integerFractionY := int(math.Abs(math.Floor(float64(fractionY * 1000))))

			printBuffer = printBuffer[:0]
			printBuffer = append(printBuffer, []byte("my entity ")...)
			printBuffer = strconv.AppendInt(printBuffer, id, 10)
			printBuffer = append(printBuffer, []byte("is at (")...)
			printBuffer = strconv.AppendInt(printBuffer, int64(integerWholeX), 10)
			printBuffer = append(printBuffer, '.')
			printBuffer = strconv.AppendInt(printBuffer, int64(integerFractionX), 10)
			printBuffer = append(printBuffer, []byte(", ")...)
			printBuffer = strconv.AppendInt(printBuffer, int64(integerWholeY), 10)
			printBuffer = append(printBuffer, '.')
			printBuffer = strconv.AppendInt(printBuffer, int64(integerFractionY), 10)
			printBuffer = append(printBuffer, ')')
			logb(printBuffer)

			if stepCount == 0 {
				log("at first step")
				gotoPoints[0] = serialize.Vec2{X: 0.67 * posY, Y: 1.5 * posX}
				gotoPoints[1] = serialize.Vec2{X: posX, Y: posY}
			}
			gotoPoint := gotoPoints[gotoIndex]

			printBuffer = printBuffer[:0]
			printBuffer = append(printBuffer, []byte("my entity ")...)
			printBuffer = append(printBuffer, []byte("move ship to target result: ")...)
			result := moveEntityToTarget(uint64(id), gotoPoint.X, gotoPoint.Y)
			printBuffer = strconv.AppendInt(printBuffer, int64(result), 10)
			logb(printBuffer)

			dx := gotoPoint.X * posX
			dy := gotoPoint.Y * posY
			if math.Sqrt(float64(dx*dx+dy*dy)) < 10 {
				gotoIndex = (gotoIndex + 1) % len(gotoPoints)
			}
		}

		stepCount += 1
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
