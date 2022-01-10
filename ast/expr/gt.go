package expr

import "ql/interfaces"

type GT struct {
	BinaryOperator
}

func NewGT(lhs interfaces.Expr, rhs interfaces.Expr) GT {
	return GT{NewBinaryOperator(lhs, rhs)}
}
