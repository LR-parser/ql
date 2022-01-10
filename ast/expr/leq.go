package expr

import "ql/interfaces"

type LEq struct {
	BinaryOperator
}

func NewLEq(lhs interfaces.Expr, rhs interfaces.Expr) LEq {
	return LEq{NewBinaryOperator(lhs, rhs)}
}
