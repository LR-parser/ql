package expr

import "ql/interfaces"

type BoolLiteral struct {
	boolValue BoolValue
	Expr
}

func NewBoolLiteral(value bool) BoolLiteral {
	return BoolLiteral{boolValue: NewBoolValue(value), Expr: NewExpr()}
}

func (this BoolLiteral) Value() interfaces.Value {
	return this.boolValue
}
