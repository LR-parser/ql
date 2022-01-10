package vari

import "ql/interfaces"

func (this VarDecl) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) {
	// store type for identifier so when we find VarExpr with this VarID we know its real type (used during typechecking)
	typeCheckArgs.Symbols().SetTypeForVarID(this.ValueType(), this.VariableIdentifier())

	// we mark it as known to indicate that earlier references to this VarID are valid
	typeCheckArgs.TypeChecker().MarkVarIDAsKnown(this.VariableIdentifier())
}
