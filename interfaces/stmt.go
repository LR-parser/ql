package interfaces

type Stmt interface {
	ASTNode
	TypeCheck(TypeCheckArgs)
}

type Form interface {
	Stmt
	Content() StmtList
	Identifier() VarID
	Questions() []Question
}

type Question interface {
	Stmt
	Label() StringLiteral
	LabelAsString() string
	VarDecl() VarDecl
	VarDeclVariableIdentifier() VarID
	VarDeclValueType() ValueType
}

type InputQuestion interface {
	Question
}

type ComputedQuestion interface {
	Question
	Computation() Expr
}

type Conditional interface {
	Stmt
	EvalConditionAsBool(VarIDValueSymbols) bool
	Condition() Expr
}

type If interface {
	Conditional
	Questions() []Question
}

type IfElse interface {
	Conditional
	IfBodyQuestions() []Question
	ElseBodyQuestions() []Question
}

type StmtList interface {
	Stmt
	Conditionals() []Conditional
	Questions() []Question
}
