package vari

import "ql/interfaces"

type VarDecl struct {
	variableIdentifier interfaces.VarID
	valueType          interfaces.ValueType
	Vari
}

func NewVarDecl(variableIdentifier interfaces.VarID, valueType interfaces.ValueType) VarDecl {
	return VarDecl{variableIdentifier: variableIdentifier, valueType: valueType, Vari: NewVari()}
}

func (this VarDecl) VariableIdentifier() interfaces.VarID {
	return this.variableIdentifier
}

func (this VarDecl) ValueType() interfaces.ValueType {
	return this.valueType
}
