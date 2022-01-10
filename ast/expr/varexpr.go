package expr

import "ql/interfaces"

type VarExpr struct {
	identifier interfaces.VarID
	Expr
}

func NewVarExpr(identifier interfaces.VarID) VarExpr {
	return VarExpr{identifier: identifier, Expr: NewExpr()}
}

func (this VarExpr) VarIdentifier() interfaces.VarID {
	return this.identifier
}
