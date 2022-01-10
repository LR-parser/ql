package symbols

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"ql/interfaces"
)

type varIDToExprSymbolTable map[interfaces.VarID]interfaces.Expr

type symbolCallBackFunction func(interfaces.VarIDValueSymbols)

type VarIDValueSymbols struct {
	Table               varIDToExprSymbolTable
	RegisteredCallbacks []func()
}

func NewVarIDValueSymbols() *VarIDValueSymbols {
	log.Info("Creating new VarIDValueSymbols")
	return &VarIDValueSymbols{Table: make(varIDToExprSymbolTable), RegisteredCallbacks: make([]func(), 0)}
}

func (this *VarIDValueSymbols) SetExprForVarID(expr interfaces.Expr, varID interfaces.VarID) {
	if expr == nil || varID == nil {
		panic("Trying to set Expr for VarID to nil or varID is nil")
	}

	this.Table[varID] = expr
	log.WithFields(log.Fields{"Identifier": varID, "Expr": expr}).Debug("Set Expr for VarID")

	for _, registeredCallback := range this.RegisteredCallbacks {
		registeredCallback()
	}
}

func (this *VarIDValueSymbols) ExprForVarID(varID interfaces.VarID) interfaces.Expr {
	if varID == nil {
		panic("Trying to get Expr for nil VarID")
	}

	expr := this.Table[varID]
	log.WithFields(log.Fields{"Identifier": varID, "Expr": expr}).Debug("Looking up Expr for VarID in SymbolTable")

	return expr
}

func (this *VarIDValueSymbols) RegisterCallback(callback func()) {
	this.RegisteredCallbacks = append(this.RegisteredCallbacks, callback)
}

/* methods related to export of symboltable to file */

type VarIDStringExprEvalSymbolTable map[string]interface{}

func (this *VarIDValueSymbols) SaveToDisk() (interface{}, error) {
	formDataAsJSON, _ := convertSymbolTableToJSON(this.convertSymbolTableKeysToStringsAndEvalValues())

	writeErr := ioutil.WriteFile("savedForm.json", formDataAsJSON, 0644)

	if writeErr != nil {
		log.WithFields(log.Fields{"error": writeErr}).Error("Encountered error during writing to disk of symbolTable data")
		return nil, writeErr
	}

	log.Info("SymbolTable written to disk as JSON")

	return formDataAsJSON, nil
}

func (this *VarIDValueSymbols) convertSymbolTableKeysToStringsAndEvalValues() VarIDStringExprEvalSymbolTable {
	symbolTableWithStringKeys := make(VarIDStringExprEvalSymbolTable)
	for varID, expr := range this.Table {
		symbolTableWithStringKeys[varID.Identifier()] = expr.Eval(this).PrimitiveValue()
	}

	return symbolTableWithStringKeys
}

func convertSymbolTableToJSON(symbolTableWithStringKeys VarIDStringExprEvalSymbolTable) ([]byte, error) {
	formDataAsJSON, jsonErr := json.MarshalIndent(symbolTableWithStringKeys, "", "  ")

	if jsonErr != nil {
		log.WithFields(log.Fields{"error": jsonErr}).Error("Encountered error during symbolTable to JSON conversion")
		return nil, jsonErr
	}

	log.WithFields(log.Fields{"formDataAsJSON": string(formDataAsJSON[:])}).Debug("Successful conversion of symbolTable to JSON")

	return formDataAsJSON, nil
}
