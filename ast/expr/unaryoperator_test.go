package expr

import (
	"github.com/stretchr/testify/assert"
	"ql/ast/vari"
	"ql/symbols"
	"testing"
)

/* Tests for unary expressions */

func TestNot(t *testing.T) {
	assert.Equal(t, NewNot(NewBoolLiteral(true)).Eval(nil), NewBoolValue(false))
}

func TestPos(t *testing.T) {
	assert.Equal(t, NewPos(NewIntegerLiteral(-10)).Eval(nil), NewIntValue(10))
}

func TestPosPos(t *testing.T) {
	assert.Equal(t, NewPos(NewPos(NewIntegerLiteral(-10))).Eval(nil), NewIntValue(10))
}

func TestNeg(t *testing.T) {
	assert.Equal(t, NewNeg(NewIntegerLiteral(10)).Eval(nil), NewIntValue(-10))
}

func TestNegNeg(t *testing.T) {
	assert.Equal(t, NewNeg(NewNeg(NewIntegerLiteral(10))).Eval(nil), NewIntValue(10))
}

func TestPosNeg(t *testing.T) {
	assert.Equal(t, NewPos(NewNeg(NewIntegerLiteral(-10))).Eval(nil), NewIntValue(10))
}

func TestNegPos(t *testing.T) {
	assert.Equal(t, NewNeg(NewPos(NewIntegerLiteral(10))).Eval(nil), NewIntValue(-10))
}

func TestVarExpr(t *testing.T) {
	exampleVarID := vari.NewVarID("TestIdentifier")

	symbols := symbols.NewVarIDValueSymbols()
	symbols.SetExprForVarID(NewIntegerLiteral(2), exampleVarID)

	assert.Equal(t, NewVarExpr(exampleVarID).Eval(symbols), NewIntValue(2))
}
