package expr

import "ql/interfaces"

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
