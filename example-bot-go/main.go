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

func (p *PrintBuffer) appendString(str string) {
	p.buffer = append(p.buffer, []byte(str)...)
}

func (p *PrintBuffer) sendAndReset() {
	logb(p.buffer)
	p.buffer = p.buffer[:0]
}

var printBuffer = &PrintBuffer{
	buffer: make([]byte, 0, 2048),
}

var stepCount uint64 = 0
var turretRotation float32

var myTeam Team = Team{
	entities: make(map[uint64]*Entity),
}
var enemyTeam Team = Team{
	entities: make(map[uint64]*Entity),
}

var gameStateDelta *gamestate.GameStateDelta = &gamestate.GameStateDelta{}

type NilU64 struct {
	IsValid bool
	Value   uint64
}

var targetEnemyId NilU64

//go:wasmexport step
func step() {
	log("step step step")

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
	gameStateDelta.Init(gameStateBuffer, n+flatbuffers.SizeUint32)

	myId := gameStateDelta.MyId()

	for idx := range gameStateDelta.FlagUpdatesLength() {
		var flag gamestate.Flag
		gameStateDelta.FlagUpdates(&flag, idx)

		if flag.OwnerId() == myId {
			myTeam.flagPos = Vec2{X: flag.X(), Y: flag.Y()}
			if stepCount == 0 {
				myTeam.basePos = myTeam.flagPos
			}
		} else {
			enemyTeam.flagPos = Vec2{X: flag.X(), Y: flag.Y()}
			if stepCount == 0 {
				enemyTeam.basePos = enemyTeam.flagPos
			}
		}
	}

	for idx := range gameStateDelta.NewEntitiesLength() {
		var entity gamestate.Entity
		gameStateDelta.NewEntities(&entity, idx)

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
		if stepCount > 0 {
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
			if len(myTeam.entities) == 0 {
				ent.target = enemyTeam.basePos
			}
			myTeam.entities[id] = ent

		} else {
			enemyTeam.entities[id] = ent
		}
	}

	for idx := range gameStateDelta.EntityUpdatesLength() {
		var entityUpdate gamestate.EntityUpdate
		gameStateDelta.EntityUpdates(&entityUpdate, idx)
		var position gamestate.Vec2
		entityUpdate.Position(&position)
		id := entityUpdate.Id()
		posX := float32(position.X())
		posY := float32(position.Y())

		_, isMine := myTeam.entities[entityUpdate.Id()]
		_, isEnemy := enemyTeam.entities[entityUpdate.Id()]

		if isMine {
			if entityUpdate.IsCommandable() {
				myTeam.entities[entityUpdate.Id()].pos = Vec2{X: posX, Y: posY}
				myTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
			} else {
				delete(myTeam.entities, id)
			}
		}

		if isEnemy {
			if entityUpdate.IsCommandable() {
				enemyTeam.entities[entityUpdate.Id()].pos = Vec2{X: posX, Y: posY}
				enemyTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
				if !targetEnemyId.IsValid {
					targetEnemyId = NilU64{
						IsValid: true,
						Value:   targetEnemyId.Value,
					}
				}
			} else {
				delete(enemyTeam.entities, id)
				if targetEnemyId.IsValid && targetEnemyId.Value == id {
					targetEnemyId.IsValid = false
				}
			}
		}
	}

	for idx := range gameStateDelta.SingleBlockEntityUpdatesLength() {
		var entityUpdate gamestate.SingleBlockEntityUpdate
		gameStateDelta.SingleBlockEntityUpdates(&entityUpdate, idx)
		var position gamestate.Vec2
		entityUpdate.Position(&position)

		_, isEnemy := myTeam.entities[entityUpdate.Id()]
		_, isMine := enemyTeam.entities[entityUpdate.Id()]

		if isMine {
			myTeam.entities[entityUpdate.Id()].pos = Vec2{X: position.X(), Y: position.Y()}
			myTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
		}

		if isEnemy {
			enemyTeam.entities[entityUpdate.Id()].pos = Vec2{X: position.X(), Y: position.Y()}
			enemyTeam.entities[entityUpdate.Id()].rotation = entityUpdate.Rotation()
		}

	}

	for idx := range gameStateDelta.DeadEntitiesLength() {
		deadEntityId := gameStateDelta.DeadEntities(idx)
		delete(enemyTeam.entities, deadEntityId)
		delete(myTeam.entities, deadEntityId)
	}

	for _, entity := range myTeam.entities {
		switch entity.entityType {
		case EntityTypeMinion:
			if targetEnemyId.IsValid {
				enemy, ok := enemyTeam.entities[targetEnemyId.Value]
				if ok {
					aimTurret(entity.id, 0, enemy.pos.X, enemy.pos.Y)
					fireCannon(entity.id, 0)
					moveEntityToTarget(entity.id, enemy.pos.X, enemy.pos.Y)
				} else {
					targetEnemyId.IsValid = false
				}
			}
		case EntityTypeShip:
			for idx := range entity.blocks {
				orientTurret(entity.id, uint32(idx), turretRotation)
				fireCannon(entity.id, uint32(idx))
				moveEntityToTarget(entity.id, entity.target.X, entity.target.Y)
				launchMissiles(entity.id, uint32(idx))
				printBuffer.appendString("moving to target (")
				printBuffer.appendFloat64(float64(entity.target.X))
				printBuffer.appendString(",")
				printBuffer.appendFloat64(float64(entity.target.Y))
				printBuffer.appendString(")")
				printBuffer.sendAndReset()
			}
		}
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
