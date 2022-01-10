package expr

import "ql/interfaces"

type Neg struct {
	UnaryOperator
}

func NewNeg(value interfaces.Expr) Neg {
	return Neg{NewUnaryOperator(value)}
}
