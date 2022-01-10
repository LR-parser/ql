package typechecker

import (
	log "github.com/Sirupsen/logrus"
	"ql/interfaces"
)

type TypeCheckArgs struct {
	typeChecker           *TypeChecker
	symbols               interfaces.TypeCheckSymbols
	currentVarDeclVisited interfaces.VarDecl
	conditionsDependentOn []interfaces.Expr
}

func NewTypeCheckArgs(typeChecker *TypeChecker, typeCheckSymbols interfaces.TypeCheckSymbols) TypeCheckArgs {
	return TypeCheckArgs{typeChecker: typeChecker, symbols: typeCheckSymbols, conditionsDependentOn: make([]interfaces.Expr, 0)}
}

func (this TypeCheckArgs) TypeChecker() interfaces.TypeChecker {
	return this.typeChecker
}

func (this TypeCheckArgs) Symbols() interfaces.TypeCheckSymbols {
	return this.symbols
}

func (this TypeCheckArgs) CurrentVarDeclVisited() interfaces.VarDecl {
	return this.currentVarDeclVisited
}

func (this TypeCheckArgs) SetCurrentVarDeclVisited(varDecl interfaces.VarDecl) interfaces.TypeCheckArgs {
	this.currentVarDeclVisited = varDecl
	return this
}

func (this TypeCheckArgs) ConditionsDependentOn() []interfaces.Expr {
	log.WithFields(log.Fields{"ConditionsDependentOn": this.conditionsDependentOn}).Info("Getting conditions dependent on")
	return this.conditionsDependentOn
}

func (this TypeCheckArgs) AddConditionDependentOn(condition interfaces.Expr) interfaces.TypeCheckArgs {
	this.conditionsDependentOn = append(this.conditionsDependentOn, condition)
	log.WithFields(log.Fields{"ConditionsDependentOn": this.conditionsDependentOn}).Info("Add dependent condition")
	return this
}
