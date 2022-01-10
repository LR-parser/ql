package expr

import "ql/interfaces"

type Pos struct {
	UnaryOperator
}

func NewPos(value interfaces.Expr) Pos {
	return Pos{NewUnaryOperator(value)}
}
