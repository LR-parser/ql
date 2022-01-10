package expr

import "ql/interfaces"

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
