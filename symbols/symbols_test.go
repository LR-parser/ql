package symbols

import (
	"io/ioutil"
	"os"
	"ql/ast/expr"
	"ql/ast/vari"
	"strings"
	"testing"
)

func TestTypeCheckSymbolsAdd(t *testing.T) {
	symbols := NewTypeCheckSymbols()
	exampleVarID := vari.NewVarID("testIdentifier")
	symbols.SetTypeForVarID(expr.NewBoolType(), exampleVarID)
	valueTypeExample := expr.NewBoolType()

	if lookupValue := symbols.TypeForVarID(exampleVarID); lookupValue != valueTypeExample {
		t.Errorf("TypeCheckSymbols not updated correctly, expected value %s for key %s, is %s", valueTypeExample, exampleVarID, lookupValue)
	}
}

func TestVarIDValueSymbolsAdd(t *testing.T) {
	symbols := NewVarIDValueSymbols()
	exampleVarID := vari.NewVarID("testIdentifier")
	exprExample := expr.NewSub(expr.NewIntegerLiteral(1), expr.NewIntegerLiteral(2))
	symbols.SetExprForVarID(exprExample, exampleVarID)

	if lookupExprValue := symbols.ExprForVarID(exampleVarID); lookupExprValue != exprExample {
		t.Errorf("VarIDValueSymbols not updated correctly, expected value %s for key %s, is %s", exprExample, exampleVarID, lookupExprValue)
	}
}

func TestVarIDValueSymbolsSaveToDisk(t *testing.T) {
	symbols := NewVarIDValueSymbols()
	exampleVarID := vari.NewVarID("testIdentifier")
	symbols.SetExprForVarID(expr.NewStringLiteral("testValue"), exampleVarID)

	symbols.SaveToDisk()

	qlFile, err := ioutil.ReadFile("savedForm.json")
	if err != nil || !strings.Contains(string(qlFile), "testIdentifier\": \"testValue") {
		t.Errorf("Output file does not contain correct data %s", qlFile)
	}

	// clean up file
	removeOutputFileAfterTest()
}

// removeOutputFileAfterTest removes the exported JSON file generated by TestVarIDValueSymbolsSaveToDisk
func removeOutputFileAfterTest() {
	err := os.Remove("savedForm.json")

	if err != nil {
		panic(err)
	}
}