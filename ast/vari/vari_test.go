package vari

import (
	"github.com/stretchr/testify/assert"
	"ql/ast/expr"
	"testing"
)

func TestNewVarID(t *testing.T) {
	varID := NewVarID("testIdentifier")

	assert.Equal(t, varID.Identifier(), "testIdentifier")
	assert.Equal(t, varID.String(), "testIdentifier")
}

func TestNewVarDecl(t *testing.T) {
	varID := NewVarID("testVarID")
	valueType := expr.NewIntegerType()
	varDecl := NewVarDecl(varID, valueType)

	assert.Equal(t, varDecl.VariableIdentifier(), varID)
	assert.Equal(t, varDecl.ValueType(), valueType)
}
