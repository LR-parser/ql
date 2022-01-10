package stmt

import (
	"github.com/stretchr/testify/assert"
	"ql/ast/expr"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/util"
	"testing"
)

func TestFormWithEmptyContent(t *testing.T) {
	identifier := vari.NewVarID("TestForm")
	exampleForm := NewForm(identifier, NewEmptyStmtList())

	assert.Equal(t, exampleForm.Identifier(), identifier)
	assert.Zero(t, len(exampleForm.Content().Questions()))
	assert.Zero(t, len(exampleForm.Content().Conditionals()))
}

func TestFormWithNonEmptyContent(t *testing.T) {
	identifier := vari.NewVarID("TestForm")
	questionExample := NewInputQuestion(expr.NewStringLiteral("What was the selling price?"), vari.NewVarDecl(vari.NewVarID("sellingPrice"), expr.NewIntegerType()))
	questionsListExample := []interfaces.Question{questionExample}
	stmtListExample := NewStmtList(questionsListExample, []interfaces.Conditional{})
	exampleForm := NewForm(identifier, stmtListExample)

	assert.Equal(t, len(exampleForm.Content().Questions()), 1)
	assert.True(t, util.AreStmtListsEqual(exampleForm.Content(), stmtListExample))
}

func TestInputQuestion(t *testing.T) {
	exampleLabel := expr.NewStringLiteral("Did you sell a house in 2010?")
	exampleVarDecl := vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType())

	exampleQuestion := NewInputQuestion(exampleLabel, exampleVarDecl)

	assert.Equal(t, exampleQuestion.Label(), exampleLabel)
}

func TestComputedQuestion(t *testing.T) {
	exampleLabel := expr.NewStringLiteral("Value residue")
	exampleVarDecl := vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewIntegerType())
	exampleComputation := expr.NewSub(expr.NewIntegerLiteral(10), expr.NewIntegerLiteral(5))

	exampleQuestion := NewComputedQuestion(exampleLabel, exampleVarDecl, exampleComputation)

	assert.Equal(t, exampleQuestion.Label(), exampleLabel)
	assert.Equal(t, exampleQuestion.Computation(), exampleComputation)
}

func TestIf(t *testing.T) {
	questionExample := NewInputQuestion(expr.NewStringLiteral("What was the selling price?"), vari.NewVarDecl(vari.NewVarID("sellingPrice"), expr.NewIntegerType()))
	ifBodyExample := NewStmtList([]interfaces.Question{questionExample}, []interfaces.Conditional{})
	ifCondExample := expr.NewBoolLiteral(true)
	ifExample := NewIf(ifCondExample, ifBodyExample)

	assert.True(t, util.AreStmtListsEqual(ifExample.Body(), ifBodyExample))
	assert.Equal(t, ifExample.Condition(), ifCondExample)
	assert.Equal(t, true, ifExample.EvalConditionAsBool(nil))
}

func TestIfElse(t *testing.T) {
	ifQuestionExample := NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	ifBodyExample := NewStmtList([]interfaces.Question{ifQuestionExample}, []interfaces.Conditional{})
	ifCondExample := expr.NewBoolLiteral(true)

	elseQuestionExample := NewInputQuestion(expr.NewStringLiteral("What was the selling price?"), vari.NewVarDecl(vari.NewVarID("sellingPrice"), expr.NewIntegerType()))
	elseBodyExample := NewStmtList([]interfaces.Question{elseQuestionExample}, []interfaces.Conditional{})

	ifElseExample := NewIfElse(ifCondExample, ifBodyExample, elseBodyExample)

	assert.True(t, util.AreStmtListsEqual(ifElseExample.IfBody(), ifBodyExample))
	assert.True(t, util.AreStmtListsEqual(ifElseExample.ElseBody(), elseBodyExample))
	assert.Equal(t, ifElseExample.Condition(), ifCondExample)
	assert.Equal(t, true, ifElseExample.EvalConditionAsBool(nil))
}

func TestStmtList(t *testing.T) {
	questionExample := NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	questionListExample := []interfaces.Question{questionExample}

	ifExample := NewIf(expr.NewBoolLiteral(true), NewEmptyStmtList())
	conditionalListExample := []interfaces.Conditional{ifExample}

	stmtListExample := NewStmtList(questionListExample, conditionalListExample)

	assert.Equal(t, len(stmtListExample.Questions()), len(questionListExample))
	assert.Equal(t, len(stmtListExample.Conditionals()), len(conditionalListExample))
}
