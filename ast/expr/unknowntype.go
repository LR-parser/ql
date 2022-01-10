package expr

import "ql/interfaces"

type UnknownType struct {
	ValueType
}

func NewUnknownType() UnknownType {
	return UnknownType{NewValueType("Unknown")}
}

func (this UnknownType) DefaultValue() interfaces.LitExpr {
	panic("UnknownType has no default value but DefaultValue method is called")

	return nil
}
