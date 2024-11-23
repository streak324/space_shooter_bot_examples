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

type Team struct {
	ownerID  byte
	basePos  Vec2
	flagPos  Vec2
	entities map[uint64]*Entity
}

type EntityType int

const (
	EntityTypeShip EntityType = iota
	EntityTypeMinion
)

type Entity struct {
	entityType EntityType
	id         uint64
	pos        Vec2
	rotation   float32
	blocks     []Block
	target     Vec2
}

type Block struct {
	featureFlags uint64
	hitpoints    float32
	isDestroyed  bool
	x            float32 // relative to ship
	y            float32 // relative to ship
	rotation     float32 // relative to ship
}

// some code taken from https://github.com/tetratelabs/wazero/blob/1e0f88bc1462ca07a33df83004914d3af7f5bcb4/examples/allocation/tinygo/testdata/greet.go

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

type PrintBuffer struct {
	buffer []byte
}

func (p *PrintBuffer) appendFloat64(data float64) {
	whole, fraction := math.Modf(data)
	integerWhole := int(math.Floor(whole))
	integerFraction := int(math.Abs(math.Floor(fraction * 1000)))

	p.buffer = strconv.AppendInt(p.buffer, int64(integerWhole), 10)
	p.buffer = append(p.buffer, '.')
	p.buffer = strconv.AppendInt(p.buffer, int64(integerFraction), 10)
}

func (p *PrintBuffer) appendInt(i int64) {
	p.buffer = strconv.AppendInt(p.buffer, i, 10)
}

func (p *PrintBuffer) appendString(str string) {
	p.buffer = append(p.buffer, []byte(str)...)
}

func (p *PrintBuffer) sendAndReset() {
	logb(p.buffer)
	p.buffer = p.buffer[:0]
}

type NilU64 struct {
	IsValid bool
	Value   uint64
}

type Globals struct {
	gameStateBuffer []byte
	printBuffer     *PrintBuffer
	stepCount       uint64
	turretRotation  float32
	myTeam          Team
	enemyTeam       Team
	gameStateDelta  *gamestate.GameStateDelta
	targetEnemyId   NilU64
}

var globals = Globals{
	gameStateBuffer: make([]byte, 64*1024),
	printBuffer: &PrintBuffer{
		buffer: make([]byte, 2048),
	},
	stepCount:      0,
	turretRotation: 0,
	myTeam: Team{
		entities: make(map[uint64]*Entity),
	},
	enemyTeam: Team{
		entities: make(map[uint64]*Entity),
	},
	gameStateDelta: &gamestate.GameStateDelta{},
	targetEnemyId:  NilU64{},
}

//go:wasmexport step
func step() {
	defer func() {
		globals.stepCount += 1
	}()

	if globals.stepCount%100 == 0 {
		globals.printBuffer.appendString("step count: ")
		globals.printBuffer.appendInt(int64(globals.stepCount))
		globals.printBuffer.appendString(" with: ")
		globals.printBuffer.appendInt(int64(len(globals.myTeam.entities)))
		globals.printBuffer.appendString(" entities")
		globals.printBuffer.sendAndReset()
	}

	globals.turretRotation = float32(globals.stepCount) * math.Pi / 60

	size := _getGameState(bytesToPtr(globals.gameStateBuffer))

	if size > int32(cap(globals.gameStateBuffer)) {
		log("error on getting game state. buffer too small")
		return
	}

	n := flatbuffers.GetUOffsetT(globals.gameStateBuffer[flatbuffers.SizeUint32:])
	globals.gameStateDelta.Init(globals.gameStateBuffer, n+flatbuffers.SizeUint32)

	myId := globals.gameStateDelta.MyId()

	for idx := range globals.gameStateDelta.FlagUpdatesLength() {
		var flag gamestate.Flag
		globals.gameStateDelta.FlagUpdates(&flag, idx)

		if flag.OwnerId() == myId {
			globals.myTeam.flagPos = Vec2{X: flag.X(), Y: flag.Y()}
			if globals.stepCount == 0 {
				globals.myTeam.basePos = globals.myTeam.flagPos
			}
		} else {
			globals.enemyTeam.flagPos = Vec2{X: flag.X(), Y: flag.Y()}
			if globals.stepCount == 0 {
				globals.enemyTeam.basePos = globals.enemyTeam.flagPos
			}
		}
	}

	for idx := range globals.gameStateDelta.NewEntitiesLength() {
		var entity gamestate.Entity
		globals.gameStateDelta.NewEntities(&entity, idx)

		if !entity.IsCommandable() {
			continue
		}

		isMine := entity.Owner() == myId
		id := entity.Id()
		var position gamestate.Vec2
		entity.Position(&position)
		posX := float32(position.X())
		posY := float32(position.Y())

		entityType := EntityTypeShip
		if globals.stepCount > 0 {
			entityType = EntityTypeMinion
		}

		blocks := make([]Block, 0, entity.BlocksLength())
		for i := 0; i < entity.BlocksLength(); i++ {
			var block gamestate.Block
			entity.Blocks(&block, i)
			blocks = append(blocks, Block{
				featureFlags: block.FeatureFlags(),
				hitpoints:    block.Hitpoints(),
				isDestroyed:  block.IsDestroyed(),
				x:            block.X(),
				y:            block.Y(),
				rotation:     block.Rotation(),
			})
		}

		ent := &Entity{
			entityType: entityType,
			id:         entity.Id(),
			pos:        Vec2{X: posX, Y: posY},
			rotation:   entity.Rotation(),
			blocks:     blocks,
		}
		if isMine {
			ent.target = Vec2{-posX, posY}
			if len(globals.myTeam.entities) == 0 {
				ent.target = globals.enemyTeam.basePos
			}
			globals.myTeam.entities[id] = ent

		} else {
			globals.enemyTeam.entities[id] = ent
		}
	}

	for idx := range globals.gameStateDelta.EntityUpdatesLength() {
		var entityUpdate gamestate.EntityUpdate
		globals.gameStateDelta.EntityUpdates(&entityUpdate, idx)
		var position gamestate.Vec2
		entityUpdate.Position(&position)
		id := entityUpdate.Id()
		posX := float32(position.X())
		posY := float32(position.Y())

		_, isMine := globals.myTeam.entities[entityUpdate.Id()]
		_, isEnemy := globals.enemyTeam.entities[entityUpdate.Id()]

		if isMine {
			if entityUpdate.IsCommandable() {
				globals.myTeam.entities[entityUpdate.Id()].pos = Vec2{X: posX, Y: posY}
				globals.myTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
			} else {
				delete(globals.myTeam.entities, id)
			}
		}

		if isEnemy {
			if entityUpdate.IsCommandable() {
				globals.enemyTeam.entities[entityUpdate.Id()].pos = Vec2{X: posX, Y: posY}
				globals.enemyTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
				if !globals.targetEnemyId.IsValid {
					globals.targetEnemyId = NilU64{
						IsValid: true,
						Value:   entityUpdate.Id(),
					}
				}
			} else {
				delete(globals.enemyTeam.entities, id)
				if globals.targetEnemyId.IsValid && globals.targetEnemyId.Value == id {
					globals.targetEnemyId.IsValid = false
				}
			}
		}
	}

	for idx := range globals.gameStateDelta.SingleBlockEntityUpdatesLength() {
		var entityUpdate gamestate.SingleBlockEntityUpdate
		globals.gameStateDelta.SingleBlockEntityUpdates(&entityUpdate, idx)
		var position gamestate.Vec2
		entityUpdate.Position(&position)

		_, isMine := globals.myTeam.entities[entityUpdate.Id()]
		_, isEnemy := globals.enemyTeam.entities[entityUpdate.Id()]

		if isMine {
			globals.myTeam.entities[entityUpdate.Id()].pos = Vec2{X: position.X(), Y: position.Y()}
			globals.myTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
		}

		if isEnemy {
			globals.enemyTeam.entities[entityUpdate.Id()].pos = Vec2{X: position.X(), Y: position.Y()}
			globals.enemyTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
		}

	}

	for idx := range globals.gameStateDelta.DeadEntitiesLength() {
		deadEntityId := globals.gameStateDelta.DeadEntities(idx)
		delete(globals.enemyTeam.entities, deadEntityId)
		delete(globals.myTeam.entities, deadEntityId)
	}

	for _, entity := range globals.myTeam.entities {
		switch entity.entityType {
		case EntityTypeMinion:
			if globals.targetEnemyId.IsValid {
				enemy, ok := globals.enemyTeam.entities[globals.targetEnemyId.Value]
				if ok {
					aimTurret(entity.id, 0, enemy.pos.X, enemy.pos.Y)
					fireCannon(entity.id, 0)
					moveEntityToTarget(entity.id, enemy.pos.X, enemy.pos.Y)
				} else {
					globals.targetEnemyId.IsValid = false
				}
			}
		case EntityTypeShip:
			for idx := range entity.blocks {
				orientTurret(entity.id, uint32(idx), globals.turretRotation)
				fireCannon(entity.id, uint32(idx))
				moveEntityToTarget(entity.id, entity.target.X, entity.target.Y)
				launchMissiles(entity.id, uint32(idx))
			}
			if globals.stepCount%30 == 0 {
				globals.printBuffer.appendString("moving to target (")
				globals.printBuffer.appendFloat64(float64(entity.target.X))
				globals.printBuffer.appendString(",")
				globals.printBuffer.appendFloat64(float64(entity.target.Y))
				globals.printBuffer.appendString(")")
				globals.printBuffer.sendAndReset()
			}
		}
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
