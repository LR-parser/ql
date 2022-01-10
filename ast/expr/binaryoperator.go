package expr

import (
	"ql/interfaces"
)

type BinaryOperator struct {
	lhs, rhs interfaces.Expr
	Expr
}

func NewBinaryOperator(lhs interfaces.Expr, rhs interfaces.Expr) BinaryOperator {
	return BinaryOperator{lhs: lhs, rhs: rhs, Expr: NewExpr()}
}

func (this BinaryOperator) LHS() interfaces.Expr {
	return this.lhs
}

func (this BinaryOperator) RHS() interfaces.Expr {
	return this.rhs
}
