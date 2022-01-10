package expr

import (
	"fmt"
	"ql/interfaces"
)

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

func (this StringValue) GEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue >= value.PrimitiveValue().(string))
}

func (this StringValue) LEq(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue <= value.PrimitiveValue().(string))
}

func (this StringValue) GT(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue > value.PrimitiveValue().(string))
}

func (this StringValue) LT(value interfaces.Value) interfaces.Value {
	return NewBoolValue(this.primitiveValue < value.PrimitiveValue().(string))
}

func (this StringValue) Add(value interfaces.Value) interfaces.Value {
	return NewStringValue(fmt.Sprintf("%s%s", this.primitiveValue, value.PrimitiveValue().(string)))
}
