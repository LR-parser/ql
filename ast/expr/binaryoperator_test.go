package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Tests for binary expression evaluation */

func TestAdd(t *testing.T) {
	assert.Equal(t, NewAdd(NewIntegerLiteral(1), NewIntegerLiteral(2)).Eval(nil), NewIntValue(3))
}

func TestMul(t *testing.T) {
	assert.Equal(t, NewMul(NewIntegerLiteral(3), NewIntegerLiteral(2)).Eval(nil), NewIntValue(6))
}

func TestMulAddPrecedence(t *testing.T) {
	assert.Equal(t, NewAdd(NewMul(NewIntegerLiteral(3), NewIntegerLiteral(2)), NewIntegerLiteral(1)).Eval(nil), NewIntValue(7))
}

func TestSub(t *testing.T) {
	assert.Equal(t, NewSub(NewIntegerLiteral(1), NewIntegerLiteral(2)).Eval(nil), NewIntValue(-1))
}

func TestDiv(t *testing.T) {
	assert.Equal(t, NewDiv(NewIntegerLiteral(9), NewIntegerLiteral(3)).Eval(nil), NewIntValue(3))
}

func TestGT(t *testing.T) {
	assert.Equal(t, NewGT(NewIntegerLiteral(3), NewIntegerLiteral(2)).Eval(nil), NewBoolValue(true))
}

func TestLT(t *testing.T) {
	assert.Equal(t, NewLT(NewIntegerLiteral(3), NewIntegerLiteral(2)).Eval(nil), NewBoolValue(false))
}

func TestGEq(t *testing.T) {
	assert.Equal(t, NewGEq(NewIntegerLiteral(3), NewIntegerLiteral(3)).Eval(nil), NewBoolValue(true))
}

func TestLEq(t *testing.T) {
	assert.Equal(t, NewLEq(NewIntegerLiteral(3), NewIntegerLiteral(3)).Eval(nil), NewBoolValue(true))
}

func TestAnd(t *testing.T) {
	assert.Equal(t, NewAnd(NewBoolLiteral(true), NewBoolLiteral(false)).Eval(nil), NewBoolValue(false))
}

func TestOr(t *testing.T) {
	assert.Equal(t, NewOr(NewBoolLiteral(true), NewBoolLiteral(false)).Eval(nil), NewBoolValue(true))
}

func TestAndOr(t *testing.T) {
	assert.Equal(t, NewAnd(NewOr(NewBoolLiteral(true), NewBoolLiteral(false)), NewBoolLiteral(false)).Eval(nil), NewBoolValue(false))
}

func TestEq(t *testing.T) {
	assert.Equal(t, NewEq(NewBoolLiteral(true), NewBoolLiteral(false)).Eval(nil), NewBoolValue(false))
}

func TestNEq(t *testing.T) {
	assert.Equal(t, NewNEq(NewBoolLiteral(true), NewBoolLiteral(false)).Eval(nil), NewBoolValue(true))
}
