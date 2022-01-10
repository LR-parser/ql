package expr

import "ql/interfaces"

type StringLiteral struct {
	stringValue StringValue
	Expr
}

func NewStringLiteral(value string) StringLiteral {
	return StringLiteral{stringValue: NewStringValue(value), Expr: NewExpr()}
}

func (this StringLiteral) Value() interfaces.Value {
	return this.stringValue
}
