package expr

import "ql/ast/node"

type Expr struct {
	node.Node
}

func NewExpr() Expr {
	return Expr{node.NewNode()}
}
