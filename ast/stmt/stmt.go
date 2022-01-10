package stmt

import "ql/ast/node"

type Stmt struct {
	node.Node
}

func NewStmt() Stmt {
	return Stmt{node.NewNode()}
}
