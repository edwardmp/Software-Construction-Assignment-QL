package expr

import (
	"math"
	"ql/interfaces"
)

type NullValue struct {
}

func NewNullValue() NullValue {
	return NullValue{}
}

func (this NullValue) PrimitiveValue() interface{} {
	return nil
}

func (this NullValue) Eq(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) NEq(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) GEq(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) LEq(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) GT(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) LT(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Div(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Mul(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Add(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Sub(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Neg() interfaces.Value {
	return nil
}

func (this NullValue) Pos() interfaces.Value {
	return nil
}

func (this NullValue) And(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Or(value interfaces.Value) interfaces.Value {
	return nil
}

func (this NullValue) Not() interfaces.Value {
	return nil
}

type IntValue struct {
	primitiveValue int
	NullValue
}

func NewIntValue(primitiveValue int) IntValue {
	return IntValue{primitiveValue: primitiveValue, NullValue: NewNullValue()}
}

func (this IntValue) PrimitiveValue() interface{} {
	return this.primitiveValue
}

func (this IntValue) Eq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue == value.PrimitiveValue().(int))
}

func (this IntValue) NEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue != value.PrimitiveValue().(int))
}

func (this IntValue) GEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue >= value.PrimitiveValue().(int))
}

func (this IntValue) LEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue >= value.PrimitiveValue().(int))
}

func (this IntValue) GT(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue > value.PrimitiveValue().(int))
}

func (this IntValue) LT(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue < value.PrimitiveValue().(int))
}

func (this IntValue) Div(value interfaces.Value) interfaces.Value {
	denominator := value.PrimitiveValue().(int)
	// since default value of int question will be set to zero, handle it here to preclude division by zero panic
	if denominator == 0 {
		return NewIntValue(0)
	}

	return NewIntValue(this.primitiveValue / denominator)
}

func (this IntValue) Mul(value interfaces.Value) interfaces.Value {
	return NewIntValue(this.primitiveValue * value.PrimitiveValue().(int))
}

func (this IntValue) Add(value interfaces.Value) interfaces.Value {
	return NewIntValue(this.primitiveValue + value.PrimitiveValue().(int))
}

func (this IntValue) Sub(value interfaces.Value) interfaces.Value {
	return NewIntValue(this.primitiveValue - value.PrimitiveValue().(int))
}

func (this IntValue) Neg() interfaces.Value {
	return NewIntValue(int(this.primitiveValue * -1))
}

func (this IntValue) Pos() interfaces.Value {
	return NewIntValue(int(math.Abs(float64(this.primitiveValue))))
}

type StringValue struct {
	primitiveValue string
	NullValue
}

func NewStringValue(primitiveValue string) StringValue {
	return StringValue{primitiveValue: primitiveValue, NullValue: NewNullValue()}
}

func (this StringValue) PrimitiveValue() interface{} {
	return this.primitiveValue
}

func (this StringValue) PrimitiveValueString() string {
	return this.primitiveValue
}

func (this StringValue) Eq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue == value.PrimitiveValue())
}

func (this StringValue) NEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue != value.PrimitiveValue())
}

type BoolValue struct {
	primitiveValue bool
	NullValue
}

func NewBoolValue(primitiveValue bool) BoolValue {
	return BoolValue{primitiveValue: primitiveValue, NullValue: NewNullValue()}
}

func (this BoolValue) PrimitiveValue() interface{} {
	return this.primitiveValue
}

func (this BoolValue) PrimitiveValueBool() bool {
	return this.primitiveValue
}

func (this BoolValue) And(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue && value.PrimitiveValue().(bool))
}

func (this BoolValue) Or(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue || value.PrimitiveValue().(bool))
}

func (this BoolValue) Not() interfaces.Value {
	return NewBoolValue(!this.primitiveValue)
}

func (this BoolValue) Eq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue == value.PrimitiveValue())
}

func (this BoolValue) NEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue != value.PrimitiveValue())
}
