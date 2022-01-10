package expr

import "ql/interfaces"

type IntegerType struct {
	ValueType
}

func NewIntegerType() IntegerType {
	return IntegerType{NewValueType("Integer")}
}

func (this IntegerType) DefaultValue() interfaces.LitExpr {
	return NewIntegerLiteral(0)
}
