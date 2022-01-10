package expr

import "ql/interfaces"

type And struct {
	BinaryOperator
}

func NewAnd(lhs interfaces.Expr, rhs interfaces.Expr) And {
	return And{NewBinaryOperator(lhs, rhs)}
}
