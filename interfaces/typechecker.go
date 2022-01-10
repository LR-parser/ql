package interfaces

type TypeChecker interface {
	AddEncounteredError(error)
	AddEncounteredWarning(error)
	EncounteredWarnings() []error
	EncounteredErrors() []error
	IsLabelUsed(StringLiteral) bool
	MarkLabelAsUsed(StringLiteral, VarDecl)
	VarIDForLabel(StringLiteral) VarID
	MarkVarIDAsKnown(VarID)
	MarkVarIDAsUnknown(VarID)
	KnownIdentifiers() map[VarID]bool
	AddDependencyForVarDecl(VarID, VarDecl)
	DependencyListForVarDeclContainsReferenceToSelf(VarDecl) bool
	ConditionsDependentOnForVarDecl(VarDecl) []Expr
}

type TypeCheckArgs interface {
	TypeChecker() TypeChecker
	Symbols() TypeCheckSymbols
	CurrentVarDeclVisited() VarDecl
	SetCurrentVarDeclVisited(VarDecl) TypeCheckArgs
	ConditionsDependentOn() []Expr
	AddConditionDependentOn(Expr) TypeCheckArgs
}
