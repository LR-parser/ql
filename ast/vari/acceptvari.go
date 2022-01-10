package vari

import "ql/interfaces"

func (this VarDecl) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitVarDecl(this, context)

	this.VariableIdentifier().Accept(visitor, context)

	return returnValue
}

func (this VarID) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	return visitor.VisitVarID(this, context)
}
