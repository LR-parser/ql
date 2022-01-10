package stmt

import (
	"ql/interfaces"
)

type StmtList struct {
	statements []interfaces.Stmt
	Stmt
}

func NewStmtList(statements []interfaces.Stmt) StmtList {
	return StmtList{statements: statements, Stmt: NewStmt()}
}

func NewEmptyStmtList() StmtList {
	return NewStmtList(nil)
}

func (this StmtList) Questions() []interfaces.Question {
	var questions []interfaces.Question
	for _, statement := range this.statements {
		if question, isQuestion := statement.(interfaces.Question); isQuestion {
			questions = append(questions, question)
		}
	}

	return questions
}

func (this StmtList) Conditionals() []interfaces.Conditional {
	var conditionals []interfaces.Conditional
	for _, statement := range this.statements {
		if conditional, isConditional := statement.(interfaces.Conditional); isConditional {
			conditionals = append(conditionals, conditional)
		}
	}

	return conditionals
}

func (this StmtList) AddStmt(questionOrConditional interfaces.Stmt) StmtList {
	this.statements = append(this.statements, questionOrConditional)

	return this
}
