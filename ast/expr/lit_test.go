package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerLiteralGetValue(t *testing.T) {
	assert.Equal(t, NewIntegerLiteral(10).Value(), NewIntValue(10))
}

func TestBoolLiteralGetValue(t *testing.T) {
	assert.Equal(t, NewBoolLiteral(true).Value(), NewBoolValue(true))
}

func TestStringLiteralGetValue(t *testing.T) {
	assert.Equal(t, NewStringLiteral("Test").Value(), NewStringValue("Test"))
}
