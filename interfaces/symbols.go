package interfaces

type TypeCheckSymbols interface {
	SetTypeForVarID(ValueType, VarID)
	TypeForVarID(VarID) ValueType
	IsTypeSetForVarID(VarID) bool
}

type VarIDValueSymbols interface {
	SetExprForVarID(Expr, VarID)
	ExprForVarID(VarID) Expr
	RegisterCallback(callback func())
	SaveToDisk() (interface{}, error)
}
