package interfaces

type ASTNode interface {
	Node
	Accept(Visitor, interface{}) interface{}
}
