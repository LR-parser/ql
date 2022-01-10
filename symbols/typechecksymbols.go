package symbols

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type varIDTovalueTypeSymbolTable map[interfaces.VarID]interfaces.ValueType

type TypeCheckSymbols struct {
	Table varIDTovalueTypeSymbolTable
}

func NewTypeCheckSymbols() *TypeCheckSymbols {
	log.Debug("Creating new TypeCheckSymbols")
	return &TypeCheckSymbols{Table: make(varIDTovalueTypeSymbolTable)}
}

func (this *TypeCheckSymbols) SetTypeForVarID(valueType interfaces.ValueType, varID interfaces.VarID) {
	if valueType == nil || varID == nil {
		panic("Trying to set value type for VarID to nil or varID is nil")
	}

	this.Table[varID] = valueType
	log.WithFields(log.Fields{"Identifier": varID, "valueType": valueType}).Debug("Set ValueType for VarID")
}

func (this *TypeCheckSymbols) TypeForVarID(varID interfaces.VarID) interfaces.ValueType {
	if varID == nil {
		panic("Trying to get type for nil VarID")
	}

	valueType := this.Table[varID]
	log.WithFields(log.Fields{"Identifier": varID, "Value": valueType}).Debug("Looking up valueType for VarID in SymbolTable")

	return valueType
}

func (this *TypeCheckSymbols) IsTypeSetForVarID(varID interfaces.VarID) bool {
	if this.TypeForVarID(varID) == nil {
		return false
	}

	return true
}
