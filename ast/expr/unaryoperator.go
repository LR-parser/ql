package expr

import "ql/interfaces"

type UnaryOperator struct {
	value interfaces.Expr
	Expr
}

func NewUnaryOperator(value interfaces.Expr) UnaryOperator {
	return UnaryOperator{value, NewExpr()}
}

func (this UnaryOperator) Value() interfaces.Expr {
	return this.value
}
