package stmt

import "ql/interfaces"

type If struct {
	condition interfaces.Expr
	body      StmtList
	Stmt
}

func NewIf(condition interfaces.Expr, body StmtList) If {
	return If{condition: condition, body: body, Stmt: NewStmt()}
}

func (this If) Body() interfaces.StmtList {
	return this.body
}

func (this If) Questions() []interfaces.Question {
	return this.body.Questions()
}

func (this If) Condition() interfaces.Expr {
	return this.condition
}

func (this If) EvalConditionAsBool(symbolTable interfaces.VarIDValueSymbols) bool {
	return this.condition.Eval(symbolTable).(interfaces.BoolValue).PrimitiveValueBool()
}
