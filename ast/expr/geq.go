package expr

import "ql/interfaces"

type GEq struct {
	BinaryOperator
}

func NewGEq(lhs interfaces.Expr, rhs interfaces.Expr) GEq {
	return GEq{NewBinaryOperator(lhs, rhs)}
}
