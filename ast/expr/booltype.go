package expr

import "ql/interfaces"

type BoolType struct {
	ValueType
}

func NewBoolType() BoolType {
	return BoolType{NewValueType("Boolean")}
}

func (this BoolType) DefaultValue() interfaces.LitExpr {
	return NewBoolLiteral(false)
}
