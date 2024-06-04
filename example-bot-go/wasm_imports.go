package main

// _log is a WebAssembly import which prints a string (linear memory offset,
// byteCount) to the console.
//
// passed in parameters is the buffer containing the string, and the byte size
//
//go:wasmimport env log
func _log(ptr, size uint32)

// _getGameState retrieves the current game state and writes it to the provided buffer.
//
// The encoded game state will be written to the buffer starting at the address
// specified by the `ptr` parameter. The `capacity` parameter specifies the
// byte capacity of the buffer. The return result is the byte size of the encoded game state.
//
// if returned byte size is larger than the buffer capacity, then consider the game state has failed to be retrieved.
// reallocate buffer size to at least as large as the returned byte size.
//
//go:wasmimport env getGameState
func _getGameState(ptr uint32, capacity uint32) int32

// command entity with id `entityId` to move to position `x`, `y`
//
//go:wasmimport env moveEntityToTarget
func moveEntityToTarget(entityId uint64, x float32, y float32) int32

// command entity with id `entityId` to rotate its block at index `blockIndex` at `rotation` radians.
//
// returns non-zero on parameter validation error
//
//go:wasmimport env orientTurret
func orientTurret(entityId uint64, blockIndex uint32, rotation float32) int32

// returns non-zero on parameter validation error
//
//go:wasmimport env fireCannon
func fireCannon(entityId uint64, blockIndex uint32) int32

// returns non-zero on parameter validation error
//
//go:wasmimport env launchMissiles
func launchMissiles(entityId uint64, blockIndex uint32) int32
