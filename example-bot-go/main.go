package main

import (
	"bot-go/gamestate"
	"math"
	"strconv"
	"unsafe"

	flatbuffers "github.com/google/flatbuffers/go"
)

type Vec2 struct {
	X float32
	Y float32
}

// some code taken from https://github.com/tetratelabs/wazero/blob/1e0f88bc1462ca07a33df83004914d3af7f5bcb4/examples/allocation/tinygo/testdata/greet.go

var printBuffer []byte = make([]byte, 2048)

var gameStateBuffer []byte = make([]byte, 64*1024)

func log(message string) {
	ptr, size := stringToPtr(message)
	_log(ptr, size)
}

func logb(message []byte) {
	ptr, size := bytesToPtr(message)
	_log(ptr, size)
}

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

type BufferTooSmall struct {
	SizeNeeded int32
}

func (b BufferTooSmall) Error() string {
	return "buffer too small"
}

func appendFloat32(dst []byte, data float64) []byte {
	whole, fraction := math.Modf(data)
	integerWhole := int(math.Floor(whole))
	integerFraction := int(math.Abs(math.Floor(fraction * 1000)))

	dst = strconv.AppendInt(dst, int64(integerWhole), 10)
	dst = append(dst, '.')
	dst = strconv.AppendInt(dst, int64(integerFraction), 10)
	return dst
}

var stepCount uint64 = 0
var gotoPoints [2]Vec2
var gotoIndex = 1
var turretRotation float32

var shipId uint64
var enemyId uint64
var enemyPosition Vec2

var gameState *gamestate.GameState = &gamestate.GameState{}

//go:export step
func step() {
	defer func() {
		stepCount += 1
	}()

	turretRotation = float32(stepCount) * math.Pi / 60

	size := _getGameState(bytesToPtr(gameStateBuffer))
	if size > int32(cap(gameStateBuffer)) {
		log("error on getting game state. buffer too small")
		return
	}

	n := flatbuffers.GetUOffsetT(gameStateBuffer[flatbuffers.SizeUint32:])
	gameState.Init(gameStateBuffer, n+flatbuffers.SizeUint32)

	for idx := range gameState.EntitiesLength() {
		var entity gamestate.Entity
		gameState.Entities(&entity, idx)
		my := entity.My()
		id := entity.Id()
		var position gamestate.Vec2
		entity.Position(&position)
		posX := float32(position.X())
		posY := float32(position.Y())

		if my {
			if stepCount == 0 {
				log("at first step")
				gotoPoints[0] = Vec2{X: 1.5 * posY, Y: 0.67 * posX}
				gotoPoints[1] = Vec2{X: posX, Y: posY}
				shipId = id
			}
			gotoPoint := gotoPoints[gotoIndex]

			dx := gotoPoint.X - posX
			dy := gotoPoint.Y - posY
			if stepCount == 0 || math.Sqrt(float64(dx*dx+dy*dy)) < 25 {
				gotoIndex = (gotoIndex + 1) % len(gotoPoints)
				printBuffer = printBuffer[:0]
				gotoPoint = gotoPoints[gotoIndex]
				printBuffer = append(printBuffer, []byte("moving to point (")...)
				printBuffer = appendFloat32(printBuffer, float64(gotoPoint.X))
				printBuffer = append(printBuffer, []byte(", ")...)
				printBuffer = appendFloat32(printBuffer, float64(gotoPoint.Y))
				printBuffer = append(printBuffer, ')')
				logb(printBuffer)
			}
			if shipId == id {
				moveEntityToTarget(id, gotoPoint.X, gotoPoint.Y)
			} else if entity.IsCommandable() {
				result := moveEntityToTarget(id, enemyPosition.X, enemyPosition.Y)
				if result != 0 {
					printBuffer = printBuffer[:0]
					printBuffer = append(printBuffer, []byte("move result ")...)
					printBuffer = strconv.AppendInt(printBuffer, int64(result), 10)
					logb(printBuffer)
				}
			}

			for blockIndex := range entity.BlocksLength() {
				var block gamestate.Block
				entity.Blocks(&block, blockIndex)
				block.FeatureFlags()
				orientTurret(entity.Id(), uint32(blockIndex), turretRotation)
				fireCannon(entity.Id(), uint32(blockIndex))
				launchMissiles(entity.Id(), uint32(blockIndex))
			}
		} else {
			if stepCount == 0 {
				enemyId = id
			}
			if enemyId == id {
				enemyPosition = Vec2{X: posX, Y: posY}
			}
		}
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
