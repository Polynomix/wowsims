package core

import (
	"time"

	"github.com/wowsims/mop/sim/core/proto"
)

type APLValueUnitIsMoving struct {
	DefaultAPLValueImpl
	unit UnitReference
}

func (rot *APLRotation) newValueUnitIsMoving(config *proto.APLValueUnitIsMoving, _ *proto.UUID) APLValue {
	unit := rot.GetSourceUnit(config.SourceUnit)
	if unit.Get() == nil {
		return nil
	}
	return &APLValueUnitIsMoving{
		unit: unit,
	}
}
func (value *APLValueUnitIsMoving) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueUnitIsMoving) GetBool(sim *Simulation) bool {
	return value.unit.Get().Moving
}
func (value *APLValueUnitIsMoving) String() string {
	return "Is Moving"
}

type APLValueUnitDistance struct {
	DefaultAPLValueImpl
	unit *Unit
}

func (rot *APLRotation) newValueUnitDistance(config *proto.APLValueUnitDistance, _ *proto.UUID) APLValue {
	return &APLValueUnitDistance{
		unit: rot.unit,
	}
}
func (value *APLValueUnitDistance) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeFloat
}
func (value *APLValueUnitDistance) GetFloat(sim *Simulation) float64 {
	return value.unit.DistanceFromTarget
}
func (value *APLValueUnitDistance) String() string {
	return "Unit Distance From Target"
}

type APLValueUnitIsHardcasting struct {
	DefaultAPLValueImpl
	unit UnitReference
}

func (rot *APLRotation) newValueUnitIsHardcasting(config *proto.APLValueUnitIsHardcasting, _ *proto.UUID) APLValue {
	unit := rot.GetSourceUnit(config.SourceUnit)
	if unit.Get() == nil {
		return nil
	}
	return &APLValueUnitIsHardcasting{
		unit: unit,
	}
}
func (value *APLValueUnitIsHardcasting) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueUnitIsHardcasting) GetBool(sim *Simulation) bool {
	return value.unit.Get().Hardcast.Expires > sim.CurrentTime
}
func (value *APLValueUnitIsHardcasting) String() string {
	return "Is Hardcasting"
}

type APLValueTimeToCastEnd struct {
	DefaultAPLValueImpl
	unit UnitReference
}

func (rot *APLRotation) newValueTimeToCastEnd(config *proto.APLValueTimeToCastEnd, _ *proto.UUID) APLValue {
	unit := rot.GetSourceUnit(config.SourceUnit)
	if unit.Get() == nil {
		return nil
	}
	return &APLValueTimeToCastEnd{
		unit: unit,
	}
}
func (value *APLValueTimeToCastEnd) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeDuration
}
func (value *APLValueTimeToCastEnd) GetDuration(sim *Simulation) time.Duration {
	return max(value.unit.Get().Hardcast.Expires-sim.CurrentTime, 0)
}
func (value *APLValueTimeToCastEnd) String() string {
	return "Time To Cast End"
}
