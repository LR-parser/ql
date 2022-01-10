package expr

import "ql/interfaces"

type LT struct {
	BinaryOperator
}

func NewLT(lhs interfaces.Expr, rhs interfaces.Expr) LT {
	return LT{NewBinaryOperator(lhs, rhs)}
}
