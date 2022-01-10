package expr

import (
	"ql/errors"
	"ql/interfaces"
)

func (this VarExpr) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	typeCheckArgs.TypeChecker().AddDependencyForVarDecl(this.VarIdentifier(), typeCheckArgs.CurrentVarDeclVisited())

	// Return the true type of the VarExpr; the type of the Expr referred to
	if typeCheckArgs.Symbols().IsTypeSetForVarID(this.VarIdentifier()) {
		return typeCheckArgs.Symbols().TypeForVarID(this.VarIdentifier()).(interfaces.ValueType)
	}

	// We don't already mark it as an error; because there is only one scope, the VarDecl may be simply declared later on
	typeCheckArgs.TypeChecker().MarkVarIDAsUnknown(this.VarIdentifier())

	// No type info in symbol table (reference to undefined question)
	return NewUnknownType()
}

func (this Add) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewIntegerType(), NewStringType()}
	actualType := this.checkOperands(allowedTypes, typeCheckArgs)

	if typeIsInExpectedTypes(actualType, allowedTypes) {
		return actualType
	}

	return NewIntegerType()
}

func (this And) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewBoolType()}
	this.checkOperands(allowedTypes, typeCheckArgs)

	return NewBoolType()
}

func (this Div) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewIntegerType()}
	this.checkOperands(allowedTypes, typeCheckArgs)

	return NewIntegerType()
}

func (this Eq) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	this.checkForEqualTypes(typeCheckArgs)
	this.checkOperands([]interfaces.ValueType{NewBoolType(), NewStringType(), NewIntegerType()}, typeCheckArgs)

	return NewBoolType()
}

func (this GEq) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	this.checkOperands([]interfaces.ValueType{NewIntegerType(), NewStringType()}, typeCheckArgs)

	return NewBoolType()
}

func (this GT) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	this.checkOperands([]interfaces.ValueType{NewIntegerType(), NewStringType()}, typeCheckArgs)
	return NewBoolType()
}

func (this LEq) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	this.checkOperands([]interfaces.ValueType{NewIntegerType(), NewStringType()}, typeCheckArgs)
	return NewBoolType()
}

func (this LT) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	this.checkOperands([]interfaces.ValueType{NewIntegerType(), NewStringType()}, typeCheckArgs)
	return NewBoolType()
}

func (this Mul) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewIntegerType()}
	this.checkOperands(allowedTypes, typeCheckArgs)
	return NewIntegerType()
}

func (this Neg) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewIntegerType()}
	this.checkOperand(allowedTypes, typeCheckArgs)
	return NewIntegerType()
}

func (this NEq) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	this.checkForEqualTypes(typeCheckArgs)
	this.checkOperands([]interfaces.ValueType{NewBoolType(), NewStringType(), NewIntegerType()}, typeCheckArgs)
	return NewBoolType()
}

func (this Not) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewBoolType()}
	this.checkOperand(allowedTypes, typeCheckArgs)
	return NewBoolType()
}

func (this Or) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewBoolType()}
	this.checkOperands(allowedTypes, typeCheckArgs)
	return NewBoolType()
}

func (this Pos) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewIntegerType()}
	this.checkOperand(allowedTypes, typeCheckArgs)
	return NewIntegerType()
}

func (this Sub) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	allowedTypes := []interfaces.ValueType{NewIntegerType()}
	this.checkOperands(allowedTypes, typeCheckArgs)
	return NewIntegerType()
}

func (this IntegerLiteral) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	return NewIntegerType()
}

func (this BoolLiteral) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	return NewBoolType()
}

func (this StringLiteral) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	return NewStringType()
}

// TypeCheck on Expr is the default implementation, which basically asserts that parent structs have overridden this method
func (this Expr) TypeCheck(typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	panic("Expr TypeCheck method not overridden")
	return nil
}

// checkOperand checks if the value of a unaryExpr is of the expected type
func (unaryExpression UnaryOperator) checkOperand(expectedTypes []interfaces.ValueType, typeCheckArgs interfaces.TypeCheckArgs) {
	checkIfOperandHasExpectedType(unaryExpression.Value(), expectedTypes, typeCheckArgs)
}

// checkOperands checks if the left-hand and right-hand sides are of the expected type
func (binaryExpr BinaryOperator) checkOperands(expectedTypes []interfaces.ValueType, typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	checkIfOperandHasExpectedType(binaryExpr.LHS(), expectedTypes, typeCheckArgs)
	return checkIfOperandHasExpectedType(binaryExpr.RHS(), expectedTypes, typeCheckArgs)
}

// checkForEqualTypes checks if the operands in a BinaryOperator have the same type, and if the types are unequal adds an error to the typechecker
func (binaryExpr BinaryOperator) checkForEqualTypes(typeCheckArgs interfaces.TypeCheckArgs) {
	lhsType := binaryExpr.LHS().TypeCheck(typeCheckArgs)
	rhsType := binaryExpr.RHS().TypeCheck(typeCheckArgs)

	if lhsType == NewUnknownType() || rhsType == NewUnknownType() {
		return
	}

	if lhsType != rhsType {
		typeCheckArgs.TypeChecker().AddEncounteredError(errors.NewOperandsOfDifferentTypesError(binaryExpr, lhsType, rhsType))
	}
}

// checkIfOperandHasExpectedType checks that an operand's actual type matches it expected type, and if not adds an error to the typechecker
func checkIfOperandHasExpectedType(expr interfaces.Expr, expectedTypes []interfaces.ValueType, typeCheckArgs interfaces.TypeCheckArgs) interfaces.ValueType {
	actualType := expr.TypeCheck(typeCheckArgs)

	if actualType == NewUnknownType() {
		return NewUnknownType()
	}

	if !typeIsInExpectedTypes(actualType, expectedTypes) {
		typeCheckArgs.TypeChecker().AddEncounteredError(errors.NewOperandWithUnexpectedTypeError(expr, expectedTypes, actualType))
	}

	return actualType
}

// typeIsInExpectedTypes returns a bool indicating of a passed type is contained in a list of passed expected types
func typeIsInExpectedTypes(typeToFind interfaces.ValueType, expectedTypes []interfaces.ValueType) bool {
	for _, expectedType := range expectedTypes {
		if expectedType == typeToFind {
			return true
		}
	}

	return false
}
