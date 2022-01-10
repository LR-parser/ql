package stmt

import "ql/interfaces"

func (this Form) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitForm(this, context)

	this.Identifier().Accept(visitor, context)
	this.Content().Accept(visitor, context)

	return returnValue
}

func (this If) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitIf(this, context)

	this.Condition().Accept(visitor, context)
	this.Body().Accept(visitor, context)

	return returnValue
}

func (this IfElse) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitIfElse(this, context)

	this.Condition().Accept(visitor, context)
	this.IfBody().Accept(visitor, context)
	this.ElseBody().Accept(visitor, context)

	return returnValue
}

func (this InputQuestion) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitInputQuestion(this, context)

	this.Label().Accept(visitor, context)
	this.VarDecl().Accept(visitor, context)

	return returnValue
}

func (this ComputedQuestion) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitComputedQuestion(this, context)

	this.Label().Accept(visitor, context)
	this.VarDecl().Accept(visitor, context)
	this.Computation().Accept(visitor, context)

	return returnValue
}

func (this StmtList) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitStmtList(this, context)

	for _, question := range this.Questions() {
		question.Accept(visitor, context)
	}

	for _, conditional := range this.Conditionals() {
		conditional.Accept(visitor, context)
	}

	return returnValue
}
