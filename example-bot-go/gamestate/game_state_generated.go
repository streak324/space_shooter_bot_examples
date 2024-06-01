// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package gamestate

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vec2 struct {
	_tab flatbuffers.Table
}

func GetRootAsVec2(buf []byte, offset flatbuffers.UOffsetT) *Vec2 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Vec2{}
	x.Init(buf, n+offset)
	return x
}

func FinishVec2Buffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsVec2(buf []byte, offset flatbuffers.UOffsetT) *Vec2 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Vec2{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedVec2Buffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Vec2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec2) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Vec2) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Vec2) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *Vec2) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Vec2) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func Vec2Start(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Vec2AddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(0, x, 0.0)
}
func Vec2AddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(1, y, 0.0)
}
func Vec2End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Entity struct {
	_tab flatbuffers.Table
}

func GetRootAsEntity(buf []byte, offset flatbuffers.UOffsetT) *Entity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Entity{}
	x.Init(buf, n+offset)
	return x
}

func FinishEntityBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsEntity(buf []byte, offset flatbuffers.UOffsetT) *Entity {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Entity{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedEntityBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Entity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Entity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Entity) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Entity) MutateId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Entity) My() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Entity) MutateMy(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *Entity) IsCommandable() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Entity) MutateIsCommandable(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *Entity) Position(obj *Vec2) *Vec2 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Vec2)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Entity) LinearVelocity(obj *Vec2) *Vec2 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Vec2)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Entity) Owner() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Entity) MutateOwner(n int32) bool {
	return rcv._tab.MutateInt32Slot(14, n)
}

func (rcv *Entity) Rotation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Entity) MutateRotation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(16, n)
}

func (rcv *Entity) AngularVelocity() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Entity) MutateAngularVelocity(n float32) bool {
	return rcv._tab.MutateFloat32Slot(18, n)
}

func EntityStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func EntityAddId(builder *flatbuffers.Builder, id uint64) {
	builder.PrependUint64Slot(0, id, 0)
}
func EntityAddMy(builder *flatbuffers.Builder, my bool) {
	builder.PrependBoolSlot(1, my, false)
}
func EntityAddIsCommandable(builder *flatbuffers.Builder, isCommandable bool) {
	builder.PrependBoolSlot(2, isCommandable, false)
}
func EntityAddPosition(builder *flatbuffers.Builder, position flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(position), 0)
}
func EntityAddLinearVelocity(builder *flatbuffers.Builder, linearVelocity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(linearVelocity), 0)
}
func EntityAddOwner(builder *flatbuffers.Builder, owner int32) {
	builder.PrependInt32Slot(5, owner, 0)
}
func EntityAddRotation(builder *flatbuffers.Builder, rotation float32) {
	builder.PrependFloat32Slot(6, rotation, 0.0)
}
func EntityAddAngularVelocity(builder *flatbuffers.Builder, angularVelocity float32) {
	builder.PrependFloat32Slot(7, angularVelocity, 0.0)
}
func EntityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type GameState struct {
	_tab flatbuffers.Table
}

func GetRootAsGameState(buf []byte, offset flatbuffers.UOffsetT) *GameState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GameState{}
	x.Init(buf, n+offset)
	return x
}

func FinishGameStateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsGameState(buf []byte, offset flatbuffers.UOffsetT) *GameState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GameState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedGameStateBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *GameState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GameState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GameState) Entities(obj *Entity, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *GameState) EntitiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func GameStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func GameStateAddEntities(builder *flatbuffers.Builder, entities flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(entities), 0)
}
func GameStateStartEntitiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GameStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
