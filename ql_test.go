package main

import (
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/lexer"
	"ql/parser"
	"ql/util"
	"testing"
)

func testStmtParse(t *testing.T, stmtAsString string, formFixture interfaces.Form) stmt.Form {
	lex := lexer.NewLexer([]byte(stmtAsString))
	parser := parser.NewParser()
	parseResult, err := parser.Parse(lex)

	if err != nil {
		t.Fatalf("Encountered fatal error during parse: %s", err)
	}

	if parsedForn, parseResultIsForm := parseResult.(stmt.Form); parseResultIsForm {
		if firstFormIdentifier, secondFormIdentifier := parsedForn.Identifier(), formFixture.Identifier(); firstFormIdentifier != secondFormIdentifier {
			t.Errorf("Form identifiers not equal: %s and %s", firstFormIdentifier, secondFormIdentifier)
		}

		if !util.AreStmtListsEqual(parsedForn.Content(), formFixture.Content()) {
			t.Errorf("Form content not equal: %v and %v", parsedForn, formFixture)
		}
	} else {
		t.Fatalf("Parse result is not form")
	}

	return parseResult.(stmt.Form)
}

/* Tests for statements */

func TestFormIdentifierExtraction(t *testing.T) {
	exampleEmptyForm := "form TestForm {}"

	exampleOutputForm := stmt.NewForm(vari.NewVarID("TestForm"), stmt.NewEmptyStmtList())
	testStmtParse(t, exampleEmptyForm, exampleOutputForm)
}

func TestFormQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean \"Did you enter a loan?\" hasMaintLoan: boolean }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	secondQuestionOutput := stmt.NewInputQuestion(expr.NewStringLiteral("Did you enter a loan?"), vari.NewVarDecl(vari.NewVarID("hasMaintLoan"), expr.NewBoolType()))
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput, secondQuestionOutput}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormComputedQuestion(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: integer \"Did you enter a loan?\" hasMaintLoan: integer \"Value residue:\" valueResidue: integer = (hasSoldHouse - hasMaintLoan) }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewIntegerType()))
	secondQuestionOutput := stmt.NewInputQuestion(expr.NewStringLiteral("Did you enter a loan?"), vari.NewVarDecl(vari.NewVarID("hasMaintLoan"), expr.NewIntegerType()))
	computedQuestion := stmt.NewComputedQuestion(expr.NewStringLiteral("Value residue:"), vari.NewVarDecl(vari.NewVarID("valueResidue"), expr.NewIntegerType()), expr.NewSub(expr.NewVarExpr(vari.NewVarID("hasSoldHouse")), expr.NewVarExpr(vari.NewVarID("hasMaintLoan"))))
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput, secondQuestionOutput, computedQuestion}, []interfaces.Conditional{})
	exampleOutputForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIf(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } }"

	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	firstQuestionBodyInput := stmt.NewInputQuestion(expr.NewStringLiteral("What was the selling price?"), vari.NewVarDecl(vari.NewVarID("sellingPrice"), expr.NewIntegerType()))
	ifBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIf(expr.NewBoolLiteral(true), ifBodyOutput)
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}

func TestFormIfElse(t *testing.T) {
	exampleFormInput := "form TestForm { \"Did you sell a house in 2010?\" hasSoldHouse: boolean if (true) { \"What was the selling price?\" sellingPrice: integer } else { \"What was the selling price?\" sellingPrice: integer } }"
	firstQuestionOutput := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	firstQuestionBodyInput := stmt.NewInputQuestion(expr.NewStringLiteral("What was the selling price?"), vari.NewVarDecl(vari.NewVarID("sellingPrice"), expr.NewIntegerType()))
	ifBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	elseBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionBodyInput}, []interfaces.Conditional{})
	ifOutput := stmt.NewIfElse(expr.NewBoolLiteral(true), ifBodyOutput, elseBodyOutput)
	exampleBodyOutput := stmt.NewStmtList([]interfaces.Question{firstQuestionOutput}, []interfaces.Conditional{ifOutput})
	exampleOutputForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBodyOutput)

	testStmtParse(t, exampleFormInput, exampleOutputForm)
}
