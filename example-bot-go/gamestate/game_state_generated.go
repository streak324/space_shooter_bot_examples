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
type MissileSlot struct {
	_tab flatbuffers.Table
}

func GetRootAsMissileSlot(buf []byte, offset flatbuffers.UOffsetT) *MissileSlot {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MissileSlot{}
	x.Init(buf, n+offset)
	return x
}

func FinishMissileSlotBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsMissileSlot(buf []byte, offset flatbuffers.UOffsetT) *MissileSlot {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MissileSlot{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedMissileSlotBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *MissileSlot) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MissileSlot) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MissileSlot) IsLoaded() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *MissileSlot) MutateIsLoaded(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *MissileSlot) ReloadTimer() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *MissileSlot) MutateReloadTimer(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *MissileSlot) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *MissileSlot) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *MissileSlot) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *MissileSlot) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *MissileSlot) Rotation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *MissileSlot) MutateRotation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func MissileSlotStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func MissileSlotAddIsLoaded(builder *flatbuffers.Builder, isLoaded bool) {
	builder.PrependBoolSlot(0, isLoaded, false)
}
func MissileSlotAddReloadTimer(builder *flatbuffers.Builder, reloadTimer float32) {
	builder.PrependFloat32Slot(1, reloadTimer, 0.0)
}
func MissileSlotAddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(2, x, 0.0)
}
func MissileSlotAddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(3, y, 0.0)
}
func MissileSlotAddRotation(builder *flatbuffers.Builder, rotation float32) {
	builder.PrependFloat32Slot(4, rotation, 0.0)
}
func MissileSlotEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Shield struct {
	_tab flatbuffers.Table
}

func GetRootAsShield(buf []byte, offset flatbuffers.UOffsetT) *Shield {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Shield{}
	x.Init(buf, n+offset)
	return x
}

func FinishShieldBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsShield(buf []byte, offset flatbuffers.UOffsetT) *Shield {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Shield{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedShieldBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Shield) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Shield) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Shield) IsDestroyed() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Shield) MutateIsDestroyed(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *Shield) Hitpoints() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Shield) MutateHitpoints(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *Shield) Radius() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Shield) MutateRadius(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func ShieldStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ShieldAddIsDestroyed(builder *flatbuffers.Builder, isDestroyed bool) {
	builder.PrependBoolSlot(0, isDestroyed, false)
}
func ShieldAddHitpoints(builder *flatbuffers.Builder, hitpoints float32) {
	builder.PrependFloat32Slot(1, hitpoints, 0.0)
}
func ShieldAddRadius(builder *flatbuffers.Builder, radius float32) {
	builder.PrependFloat32Slot(2, radius, 0.0)
}
func ShieldEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Block struct {
	_tab flatbuffers.Table
}

func GetRootAsBlock(buf []byte, offset flatbuffers.UOffsetT) *Block {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Block{}
	x.Init(buf, n+offset)
	return x
}

func FinishBlockBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsBlock(buf []byte, offset flatbuffers.UOffsetT) *Block {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Block{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedBlockBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Block) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Block) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Block) BlockType() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Block) MutateBlockType(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Block) FeatureFlags() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Block) MutateFeatureFlags(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *Block) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Block) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *Block) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Block) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *Block) Rotation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Block) MutateRotation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func (rcv *Block) Hitpoints() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Block) MutateHitpoints(n float32) bool {
	return rcv._tab.MutateFloat32Slot(14, n)
}

func (rcv *Block) AppliedThrust() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Block) MutateAppliedThrust(n float32) bool {
	return rcv._tab.MutateFloat32Slot(16, n)
}

func (rcv *Block) IsDestroyed() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Block) MutateIsDestroyed(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *Block) MissileSlots(obj *MissileSlot, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Block) MissileSlotsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Block) LocalTurretRotation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Block) MutateLocalTurretRotation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(22, n)
}

func (rcv *Block) Shield(obj *Shield) *Shield {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Shield)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func BlockStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}
func BlockAddBlockType(builder *flatbuffers.Builder, blockType uint32) {
	builder.PrependUint32Slot(0, blockType, 0)
}
func BlockAddFeatureFlags(builder *flatbuffers.Builder, featureFlags uint64) {
	builder.PrependUint64Slot(1, featureFlags, 0)
}
func BlockAddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(2, x, 0.0)
}
func BlockAddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(3, y, 0.0)
}
func BlockAddRotation(builder *flatbuffers.Builder, rotation float32) {
	builder.PrependFloat32Slot(4, rotation, 0.0)
}
func BlockAddHitpoints(builder *flatbuffers.Builder, hitpoints float32) {
	builder.PrependFloat32Slot(5, hitpoints, 0.0)
}
func BlockAddAppliedThrust(builder *flatbuffers.Builder, appliedThrust float32) {
	builder.PrependFloat32Slot(6, appliedThrust, 0.0)
}
func BlockAddIsDestroyed(builder *flatbuffers.Builder, isDestroyed bool) {
	builder.PrependBoolSlot(7, isDestroyed, false)
}
func BlockAddMissileSlots(builder *flatbuffers.Builder, missileSlots flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(missileSlots), 0)
}
func BlockStartMissileSlotsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BlockAddLocalTurretRotation(builder *flatbuffers.Builder, localTurretRotation float32) {
	builder.PrependFloat32Slot(9, localTurretRotation, 0.0)
}
func BlockAddShield(builder *flatbuffers.Builder, shield flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(shield), 0)
}
func BlockEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Explosion struct {
	_tab flatbuffers.Table
}

func GetRootAsExplosion(buf []byte, offset flatbuffers.UOffsetT) *Explosion {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Explosion{}
	x.Init(buf, n+offset)
	return x
}

func FinishExplosionBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsExplosion(buf []byte, offset flatbuffers.UOffsetT) *Explosion {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Explosion{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedExplosionBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Explosion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Explosion) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Explosion) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Explosion) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *Explosion) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Explosion) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *Explosion) Radius() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Explosion) MutateRadius(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *Explosion) Damage() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Explosion) MutateDamage(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func ExplosionStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ExplosionAddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(0, x, 0.0)
}
func ExplosionAddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(1, y, 0.0)
}
func ExplosionAddRadius(builder *flatbuffers.Builder, radius float32) {
	builder.PrependFloat32Slot(2, radius, 0.0)
}
func ExplosionAddDamage(builder *flatbuffers.Builder, damage float32) {
	builder.PrependFloat32Slot(3, damage, 0.0)
}
func ExplosionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
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

func (rcv *Entity) Owner() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Entity) MutateOwner(n byte) bool {
	return rcv._tab.MutateByteSlot(14, n)
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

func (rcv *Entity) Blocks(obj *Block, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Entity) BlocksLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EntityStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
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
func EntityAddOwner(builder *flatbuffers.Builder, owner byte) {
	builder.PrependByteSlot(5, owner, 0)
}
func EntityAddRotation(builder *flatbuffers.Builder, rotation float32) {
	builder.PrependFloat32Slot(6, rotation, 0.0)
}
func EntityAddAngularVelocity(builder *flatbuffers.Builder, angularVelocity float32) {
	builder.PrependFloat32Slot(7, angularVelocity, 0.0)
}
func EntityAddBlocks(builder *flatbuffers.Builder, blocks flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(blocks), 0)
}
func EntityStartBlocksVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EntityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Projectile struct {
	_tab flatbuffers.Table
}

func GetRootAsProjectile(buf []byte, offset flatbuffers.UOffsetT) *Projectile {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Projectile{}
	x.Init(buf, n+offset)
	return x
}

func FinishProjectileBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsProjectile(buf []byte, offset flatbuffers.UOffsetT) *Projectile {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Projectile{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedProjectileBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Projectile) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Projectile) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Projectile) My() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Projectile) MutateMy(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *Projectile) Position(obj *Vec2) *Vec2 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
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

func (rcv *Projectile) LinearVelocity(obj *Vec2) *Vec2 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
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

func (rcv *Projectile) Size() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Projectile) MutateSize(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *Projectile) Damage() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Projectile) MutateDamage(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func ProjectileStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ProjectileAddMy(builder *flatbuffers.Builder, my bool) {
	builder.PrependBoolSlot(0, my, false)
}
func ProjectileAddPosition(builder *flatbuffers.Builder, position flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(position), 0)
}
func ProjectileAddLinearVelocity(builder *flatbuffers.Builder, linearVelocity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(linearVelocity), 0)
}
func ProjectileAddSize(builder *flatbuffers.Builder, size float32) {
	builder.PrependFloat32Slot(3, size, 0.0)
}
func ProjectileAddDamage(builder *flatbuffers.Builder, damage float32) {
	builder.PrependFloat32Slot(4, damage, 0.0)
}
func ProjectileEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Flag struct {
	_tab flatbuffers.Table
}

func GetRootAsFlag(buf []byte, offset flatbuffers.UOffsetT) *Flag {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Flag{}
	x.Init(buf, n+offset)
	return x
}

func FinishFlagBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsFlag(buf []byte, offset flatbuffers.UOffsetT) *Flag {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Flag{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedFlagBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Flag) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Flag) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Flag) OwnerId() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Flag) MutateOwnerId(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *Flag) X() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Flag) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *Flag) Y() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Flag) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *Flag) CarrierId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Flag) MutateCarrierId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func FlagStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FlagAddOwnerId(builder *flatbuffers.Builder, ownerId byte) {
	builder.PrependByteSlot(0, ownerId, 0)
}
func FlagAddX(builder *flatbuffers.Builder, x float32) {
	builder.PrependFloat32Slot(1, x, 0.0)
}
func FlagAddY(builder *flatbuffers.Builder, y float32) {
	builder.PrependFloat32Slot(2, y, 0.0)
}
func FlagAddCarrierId(builder *flatbuffers.Builder, carrierId uint64) {
	builder.PrependUint64Slot(3, carrierId, 0)
}
func FlagEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
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

func (rcv *GameState) Flags(obj *Flag, j int) bool {
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

func (rcv *GameState) FlagsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GameState) Entities(obj *Entity, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
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
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GameState) Projectiles(obj *Projectile, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *GameState) ProjectilesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GameState) Explosions(obj *Explosion, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *GameState) ExplosionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *GameState) MyId() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GameState) MutateMyId(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func (rcv *GameState) WinnerId() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GameState) MutateWinnerId(n byte) bool {
	return rcv._tab.MutateByteSlot(14, n)
}

func GameStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func GameStateAddFlags(builder *flatbuffers.Builder, flags flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(flags), 0)
}
func GameStateStartFlagsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GameStateAddEntities(builder *flatbuffers.Builder, entities flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(entities), 0)
}
func GameStateStartEntitiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GameStateAddProjectiles(builder *flatbuffers.Builder, projectiles flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(projectiles), 0)
}
func GameStateStartProjectilesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GameStateAddExplosions(builder *flatbuffers.Builder, explosions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(explosions), 0)
}
func GameStateStartExplosionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GameStateAddMyId(builder *flatbuffers.Builder, myId byte) {
	builder.PrependByteSlot(4, myId, 0)
}
func GameStateAddWinnerId(builder *flatbuffers.Builder, winnerId byte) {
	builder.PrependByteSlot(5, winnerId, 0)
}
func GameStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
