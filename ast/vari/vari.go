package vari

import "ql/ast/node"

type Vari struct {
	node.Node
}

func NewVari() Vari {
	return Vari{node.NewNode()}
}
