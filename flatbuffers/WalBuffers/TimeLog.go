// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package WalBuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TimeLog struct {
	_tab flatbuffers.Table
}

func GetRootAsTimeLog(buf []byte, offset flatbuffers.UOffsetT) *TimeLog {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TimeLog{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TimeLog) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TimeLog) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TimeLog) TimeId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TimeLog) MutateTimeId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *TimeLog) X() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 99999999
}

func (rcv *TimeLog) MutateX(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *TimeLog) Y() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 99999999
}

func (rcv *TimeLog) MutateY(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *TimeLog) Angle() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 99999999.0
}

func (rcv *TimeLog) MutateAngle(n float32) bool {
	return rcv._tab.MutateFloat32Slot(10, n)
}

func (rcv *TimeLog) CannonAngle() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 99999999.0
}

func (rcv *TimeLog) MutateCannonAngle(n float32) bool {
	return rcv._tab.MutateFloat32Slot(12, n)
}

func (rcv *TimeLog) CannonRotation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 99999999.0
}

func (rcv *TimeLog) MutateCannonRotation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(14, n)
}

func (rcv *TimeLog) CannonUntilTimeId() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 99999999
}

func (rcv *TimeLog) MutateCannonUntilTimeId(n int64) bool {
	return rcv._tab.MutateInt64Slot(16, n)
}

func (rcv *TimeLog) Fire() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *TimeLog) MutateFire(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *TimeLog) IsDelete() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *TimeLog) MutateIsDelete(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

func (rcv *TimeLog) Explode() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *TimeLog) MutateExplode(n bool) bool {
	return rcv._tab.MutateBoolSlot(22, n)
}

func (rcv *TimeLog) ExplodeOther() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *TimeLog) MutateExplodeOther(n bool) bool {
	return rcv._tab.MutateBoolSlot(24, n)
}

func (rcv *TimeLog) DeleteOtherIds(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TimeLog) DeleteOtherIdsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TimeLog) MutateDeleteOtherIds(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *TimeLog) VelocityX() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 99999999.0
}

func (rcv *TimeLog) MutateVelocityX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(28, n)
}

func (rcv *TimeLog) VelocityY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 99999999.0
}

func (rcv *TimeLog) MutateVelocityY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(30, n)
}

func (rcv *TimeLog) VelocityRotation() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 99999999.0
}

func (rcv *TimeLog) MutateVelocityRotation(n float32) bool {
	return rcv._tab.MutateFloat32Slot(32, n)
}

func (rcv *TimeLog) VelocityUntilTimeId() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 99999999
}

func (rcv *TimeLog) MutateVelocityUntilTimeId(n int64) bool {
	return rcv._tab.MutateInt64Slot(34, n)
}

func TimeLogStart(builder *flatbuffers.Builder) {
	builder.StartObject(16)
}
func TimeLogAddTimeId(builder *flatbuffers.Builder, timeId uint64) {
	builder.PrependUint64Slot(0, timeId, 0)
}
func TimeLogAddX(builder *flatbuffers.Builder, x int32) {
	builder.PrependInt32Slot(1, x, 99999999)
}
func TimeLogAddY(builder *flatbuffers.Builder, y int32) {
	builder.PrependInt32Slot(2, y, 99999999)
}
func TimeLogAddAngle(builder *flatbuffers.Builder, angle float32) {
	builder.PrependFloat32Slot(3, angle, 99999999.0)
}
func TimeLogAddCannonAngle(builder *flatbuffers.Builder, cannonAngle float32) {
	builder.PrependFloat32Slot(4, cannonAngle, 99999999.0)
}
func TimeLogAddCannonRotation(builder *flatbuffers.Builder, cannonRotation float32) {
	builder.PrependFloat32Slot(5, cannonRotation, 99999999.0)
}
func TimeLogAddCannonUntilTimeId(builder *flatbuffers.Builder, cannonUntilTimeId int64) {
	builder.PrependInt64Slot(6, cannonUntilTimeId, 99999999)
}
func TimeLogAddFire(builder *flatbuffers.Builder, fire bool) {
	builder.PrependBoolSlot(7, fire, false)
}
func TimeLogAddIsDelete(builder *flatbuffers.Builder, isDelete bool) {
	builder.PrependBoolSlot(8, isDelete, false)
}
func TimeLogAddExplode(builder *flatbuffers.Builder, explode bool) {
	builder.PrependBoolSlot(9, explode, false)
}
func TimeLogAddExplodeOther(builder *flatbuffers.Builder, explodeOther bool) {
	builder.PrependBoolSlot(10, explodeOther, false)
}
func TimeLogAddDeleteOtherIds(builder *flatbuffers.Builder, deleteOtherIds flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(deleteOtherIds), 0)
}
func TimeLogStartDeleteOtherIdsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TimeLogAddVelocityX(builder *flatbuffers.Builder, velocityX float32) {
	builder.PrependFloat32Slot(12, velocityX, 99999999.0)
}
func TimeLogAddVelocityY(builder *flatbuffers.Builder, velocityY float32) {
	builder.PrependFloat32Slot(13, velocityY, 99999999.0)
}
func TimeLogAddVelocityRotation(builder *flatbuffers.Builder, velocityRotation float32) {
	builder.PrependFloat32Slot(14, velocityRotation, 99999999.0)
}
func TimeLogAddVelocityUntilTimeId(builder *flatbuffers.Builder, velocityUntilTimeId int64) {
	builder.PrependInt64Slot(15, velocityUntilTimeId, 99999999)
}
func TimeLogEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
