package expr

import "ql/interfaces"

type IntegerLiteral struct {
	intValue IntValue
	Expr
}

func NewIntegerLiteral(value int) IntegerLiteral {
	return IntegerLiteral{intValue: NewIntValue(value), Expr: NewExpr()}
}

func (this IntegerLiteral) Value() interfaces.Value {
	return this.intValue
}
