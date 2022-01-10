package expr

import "ql/interfaces"

type StringType struct {
	ValueType
}

func NewStringType() StringType {
	return StringType{NewValueType("String")}
}

func (this StringType) DefaultValue() interfaces.LitExpr {
	return NewStringLiteral("")
}
