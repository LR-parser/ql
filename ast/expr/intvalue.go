package expr

import (
	"math"
	"ql/interfaces"
)

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
