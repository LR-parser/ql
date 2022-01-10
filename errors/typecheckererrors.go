package errors

import (
	"fmt"
	"ql/interfaces"
)

const (
	undefinedQuestionReferenceErrorString           = "Reference to undefined question found"
	nonBooleanConditionErrorString                  = "Non-boolean type used as condition"
	cyclicDependencyErrorString                     = "Cyclic dependency encountered"
	duplicateLabelWarningString                     = "Duplicate label detected"
	questionRedeclaredWithDifferentTypesErrorString = "Question redeclared with different types"
	declaratedTypeAndActualTypeDeviateErrorString   = "Encountered computed question with mismatch between declared type and actual computation type"
	operandsOfDifferentTypesErrorString             = "Encountered BinaryOperator with operands of different types"
	operandWithUnexpectedTypeErrorString            = "Encountered operand with unexpected type"
)

type TypeCheckError struct {
	error
}

type UndefinedQuestionReferenceError struct {
	TypeCheckError
	UndefinedVarID interfaces.VarID
}

func NewUndefinedQuestionReferenceError(undefinedVarID interfaces.VarID) UndefinedQuestionReferenceError {
	return UndefinedQuestionReferenceError{UndefinedVarID: undefinedVarID, TypeCheckError: TypeCheckError{fmt.Errorf(undefinedQuestionReferenceErrorString)}}
}

func (this UndefinedQuestionReferenceError) Error() string {
	return fmt.Sprintf("%s: %s", this.TypeCheckError.Error(), this.UndefinedVarID)
}

type NonBooleanConditionError struct {
	TypeCheckError
	ViolatingCondition       interfaces.Expr
	TypeOfViolatingCondition interfaces.ValueType
}

func NewNonBooleanConditionError(violatingCondition interfaces.Expr, typeOfViolatingCondition interfaces.ValueType) NonBooleanConditionError {
	return NonBooleanConditionError{ViolatingCondition: violatingCondition, TypeOfViolatingCondition: typeOfViolatingCondition, TypeCheckError: TypeCheckError{fmt.Errorf(nonBooleanConditionErrorString)}}
}

func (this NonBooleanConditionError) Error() string {
	return fmt.Sprintf("%s: %s is of type %s", this.TypeCheckError.Error(), this.ViolatingCondition, this.TypeOfViolatingCondition)
}

type CyclicDependencyError struct {
	TypeCheckError
	VarDeclWithCyclicDependency interfaces.VarDecl
}

func NewCyclicDependencyError(varDeclWithCyclicDependency interfaces.VarDecl) CyclicDependencyError {
	return CyclicDependencyError{VarDeclWithCyclicDependency: varDeclWithCyclicDependency, TypeCheckError: TypeCheckError{fmt.Errorf(cyclicDependencyErrorString)}}
}

func (this CyclicDependencyError) Error() string {
	return fmt.Sprintf("%s: %s eventually refers to self", this.TypeCheckError.Error(), this.VarDeclWithCyclicDependency.VariableIdentifier())
}

type DuplicateLabelWarning struct {
	TypeCheckError
	OriginalQuestion         interfaces.Question
	ConflictingQuestionVarID interfaces.VarID
}

func NewDuplicateLabelWarning(originalQuestion interfaces.Question, conflictingQuestionVarID interfaces.VarID) DuplicateLabelWarning {
	return DuplicateLabelWarning{OriginalQuestion: originalQuestion, ConflictingQuestionVarID: conflictingQuestionVarID, TypeCheckError: TypeCheckError{fmt.Errorf(duplicateLabelWarningString)}}
}

func (this DuplicateLabelWarning) Error() string {
	return fmt.Sprintf("%s: label %s already used for question with identifier %s, using again for question with identifier %s", this.TypeCheckError.Error(), this.OriginalQuestion.Label(), this.ConflictingQuestionVarID, this.OriginalQuestion.VarDeclVariableIdentifier())
}

type QuestionRedeclaredWithDifferentTypesError struct {
	TypeCheckError
	TypeOfFirstOccurence  interfaces.ValueType
	TypeOfSecondOccurence interfaces.ValueType
}

func NewQuestionRedeclaredWithDifferentTypesError(typeFirstOccurence interfaces.ValueType, typeSecondOccurence interfaces.ValueType) QuestionRedeclaredWithDifferentTypesError {
	return QuestionRedeclaredWithDifferentTypesError{TypeOfFirstOccurence: typeFirstOccurence, TypeOfSecondOccurence: typeSecondOccurence, TypeCheckError: TypeCheckError{fmt.Errorf(questionRedeclaredWithDifferentTypesErrorString)}}
}

func (this QuestionRedeclaredWithDifferentTypesError) Error() string {
	return fmt.Sprintf("%s: %s and %s", this.TypeCheckError.Error(), this.TypeOfFirstOccurence, this.TypeOfSecondOccurence)
}

type DeclaratedTypeAndActualTypeDeviateError struct {
	TypeCheckError
	ExpectedType interfaces.ValueType
	ActualType   interfaces.ValueType
}

func NewDeclaratedTypeAndActualTypeDeviateError(expectedType interfaces.ValueType, actualType interfaces.ValueType) DeclaratedTypeAndActualTypeDeviateError {
	return DeclaratedTypeAndActualTypeDeviateError{ExpectedType: expectedType, ActualType: actualType, TypeCheckError: TypeCheckError{fmt.Errorf(declaratedTypeAndActualTypeDeviateErrorString)}}
}

func (this DeclaratedTypeAndActualTypeDeviateError) Error() string {
	return fmt.Sprintf("%s: expected type %s but actually is of type %s", this.TypeCheckError.Error(), this.ExpectedType, this.ActualType)
}

type OperandsOfDifferentTypesError struct {
	TypeCheckError
	FirstOperandType   interfaces.ValueType
	SecondOperandType  interfaces.ValueType
	AffectedExpression interfaces.Expr
}

func NewOperandsOfDifferentTypesError(affectedExpression interfaces.Expr, firstOperandType interfaces.ValueType, secondOperandType interfaces.ValueType) OperandsOfDifferentTypesError {
	return OperandsOfDifferentTypesError{AffectedExpression: affectedExpression, FirstOperandType: firstOperandType, SecondOperandType: secondOperandType, TypeCheckError: TypeCheckError{fmt.Errorf(operandsOfDifferentTypesErrorString)}}
}

func (this OperandsOfDifferentTypesError) Error() string {
	return fmt.Sprintf("%s: type of first operand is %s but type of second operand is %s in expression %s", this.TypeCheckError.Error(), this.FirstOperandType, this.SecondOperandType, this.AffectedExpression)
}

type OperandWithUnexpectedTypeError struct {
	TypeCheckError
	ExpectedTypes      []interfaces.ValueType
	ActualType         interfaces.ValueType
	AffectedExpression interfaces.Expr
}

func NewOperandWithUnexpectedTypeError(affectedExpression interfaces.Expr, expectedTypes []interfaces.ValueType, actualType interfaces.ValueType) OperandWithUnexpectedTypeError {
	return OperandWithUnexpectedTypeError{AffectedExpression: affectedExpression, ExpectedTypes: expectedTypes, ActualType: actualType, TypeCheckError: TypeCheckError{fmt.Errorf(operandWithUnexpectedTypeErrorString)}}
}

func (this OperandWithUnexpectedTypeError) Error() string {
	return fmt.Sprintf("%s: expected type of operand was one of %s but actual type is %s in expression %s", this.TypeCheckError.Error(), this.ExpectedTypes, this.ActualType, this.AffectedExpression)
}
