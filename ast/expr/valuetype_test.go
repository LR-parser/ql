package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVarTypeDefaultValueForIntegerType(t *testing.T) {
	varTypeExample := NewIntegerType()
	correctDefaultValue := NewIntegerLiteral(0)
	assert.Equal(t, varTypeExample.DefaultValue(), correctDefaultValue)
}

func TestVarTypeDefaultValueForBoolType(t *testing.T) {
	varTypeExample := NewBoolType()
	correctDefaultValue := NewBoolLiteral(false)
	assert.Equal(t, varTypeExample.DefaultValue(), correctDefaultValue)
}

func TestVarTypeDefaultValueForStringType(t *testing.T) {
	varTypeExample := NewStringType()
	correctDefaultValue := NewStringLiteral("")
	assert.Equal(t, varTypeExample.DefaultValue(), correctDefaultValue)
}
