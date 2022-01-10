package stmt

import "ql/interfaces"

type ComputedQuestion struct {
	computation interfaces.Expr
	Question
}

func NewComputedQuestion(label interfaces.StringLiteral, varDecl interfaces.VarDecl, computation interfaces.Expr) ComputedQuestion {
	return ComputedQuestion{computation: computation, Question: NewQuestion(label, varDecl)}
}

func (this ComputedQuestion) Computation() interfaces.Expr {
	return this.computation
}
