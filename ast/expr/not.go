package expr

import "ql/interfaces"

type Not struct {
	UnaryOperator
}

func NewNot(value interfaces.Expr) Not {
	return Not{NewUnaryOperator(value)}
}
