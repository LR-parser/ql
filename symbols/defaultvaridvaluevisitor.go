package symbols

import (
	"ql/ast/visitor"
	"ql/interfaces"
)

type DefaultVarIDValueVisitor struct {
	*visitor.BaseVisitor
}

func NewDefaultVarIDValueVisitor() *DefaultVarIDValueVisitor {
	return &DefaultVarIDValueVisitor{visitor.NewBaseVisitor()}
}

func (this *DefaultVarIDValueVisitor) StartSettingDefaultValuesForVarIDs(form interfaces.Form) *VarIDValueSymbols {
	varIDValueSymbols := NewVarIDValueSymbols()
	form.Accept(this, varIDValueSymbols)

	return varIDValueSymbols
}

func (this *DefaultVarIDValueVisitor) VisitVarDecl(varDecl interfaces.VarDecl, context interface{}) interface{} {
	symbols := context.(interfaces.VarIDValueSymbols)
	symbols.SetExprForVarID(varDecl.ValueType().DefaultValue(), varDecl.VariableIdentifier())

	return nil
}
