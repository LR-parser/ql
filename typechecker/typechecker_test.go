package typechecker

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/errors"
	"ql/interfaces"
	"ql/symbols"
	"testing"
)

type TypeCheckerTestSuite struct {
	suite.Suite
	typeCheckArgs TypeCheckArgs
}

func (suite *TypeCheckerTestSuite) SetupTest() {
	typeChecker := NewTypeChecker()
	suite.typeCheckArgs = NewTypeCheckArgs(typeChecker, symbols.NewTypeCheckSymbols())
}

/* helper methods */
func (suite *TypeCheckerTestSuite) testThatNumberOfErrorsOfPassedTypeArePresent(expectedType interface{}, errorsReported []error, expectedAmountOfErrors int) {
	if assert.Equal(suite.T(), len(errorsReported), expectedAmountOfErrors) {
		for _, reportedError := range errorsReported {
			assert.IsType(suite.T(), expectedType, reportedError)
		}
	}
}

func (suite *TypeCheckerTestSuite) testThatNumberOfErrorsOfPassedTypeArePresentForExpr(expr interfaces.Expr, expectedType interface{}, expectedAmountOfErrors int) {
	errorsReported := suite.typeCheckPassedExprAndReturnErrors(expr)
	suite.testThatNumberOfErrorsOfPassedTypeArePresent(expectedType, errorsReported, expectedAmountOfErrors)
}

func (suite *TypeCheckerTestSuite) testThatNumberOfErrorsOfPassedTypeArePresentForForm(form interfaces.Form, expectedType interface{}, expectedAmountOfErrors int) {
	errorsReported, _ := suite.typeCheckPassedFormAndReturnErrors(form)
	suite.testThatNumberOfErrorsOfPassedTypeArePresent(expectedType, errorsReported, expectedAmountOfErrors)
}

func (suite *TypeCheckerTestSuite) testThatNumberOfWarningsOfPassedTypeArePresentForForm(form interfaces.Form, expectedType interface{}, expectedAmountOfErrors int) {
	_, warningsReported := suite.typeCheckPassedFormAndReturnErrors(form)
	suite.testThatNumberOfErrorsOfPassedTypeArePresent(expectedType, warningsReported, expectedAmountOfErrors)
}

func (suite *TypeCheckerTestSuite) typeCheckPassedExprAndReturnErrors(expr interfaces.Expr) []error {
	expr.TypeCheck(suite.typeCheckArgs)
	return suite.typeCheckArgs.TypeChecker().EncounteredErrors()
}

func (suite *TypeCheckerTestSuite) typeCheckPassedFormAndReturnErrors(form interfaces.Form) ([]error, []error) {
	form.TypeCheck(suite.typeCheckArgs)
	return suite.typeCheckArgs.TypeChecker().EncounteredErrors(), suite.typeCheckArgs.TypeChecker().EncounteredWarnings()
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TypeCheckerTestSuite))
}

/* actual tests */

func (suite *TypeCheckerTestSuite) TestInvalidOperandsCheckerForDifferentOperandEvalTypes() {
	exampleExpr := expr.NewSub(expr.NewBoolLiteral(true), expr.NewIntegerLiteral(10))

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForExpr(exampleExpr, errors.OperandWithUnexpectedTypeError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestInvalidOperandsCheckerForInvalidBinaryOperationWithIntegers() {
	exampleExpr := expr.NewAnd(expr.NewIntegerLiteral(10), expr.NewIntegerLiteral(8))
	suite.testThatNumberOfErrorsOfPassedTypeArePresentForExpr(exampleExpr, errors.OperandWithUnexpectedTypeError{}, 2)
}

func (suite *TypeCheckerTestSuite) TestInvalidOperandsCheckerForInvalidBinaryOperationWithStrings() {
	exampleExpr := expr.NewAnd(expr.NewStringLiteral("Test A"), expr.NewStringLiteral("Test B"))
	suite.testThatNumberOfErrorsOfPassedTypeArePresentForExpr(exampleExpr, errors.OperandWithUnexpectedTypeError{}, 2)
}

func (suite *TypeCheckerTestSuite) TestInvalidOperandsCheckerForInvalidUnaryOperationWithBool() {
	exampleExpr := expr.NewNeg(expr.NewBoolLiteral(true))
	suite.testThatNumberOfErrorsOfPassedTypeArePresentForExpr(exampleExpr, errors.OperandWithUnexpectedTypeError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestInvalidOperandsCheckerForInvalidUnaryOperationWithInt() {
	exampleExpr := expr.NewNeg(expr.NewBoolLiteral(true))
	suite.testThatNumberOfErrorsOfPassedTypeArePresentForExpr(exampleExpr, errors.OperandWithUnexpectedTypeError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestInvalidOperandsCheckerForInvalidUnaryOperationWithString() {
	exampleExpr := expr.NewNot(expr.NewStringLiteral("Test"))
	suite.testThatNumberOfErrorsOfPassedTypeArePresentForExpr(exampleExpr, errors.OperandWithUnexpectedTypeError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestUndefinedQuestionReferenceChecker() {
	computedQuestion := stmt.NewComputedQuestion(expr.NewStringLiteral("Value residue:"), vari.NewVarDecl(vari.NewVarID("valueResidue"), expr.NewIntegerType()), expr.NewSub(expr.NewIntegerLiteral(100), expr.NewVarExpr(vari.NewVarID("hasMaintLoan"))))
	exampleBody := stmt.NewStmtList([]interfaces.Question{computedQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.UndefinedQuestionReferenceError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestNonBoolConditionalChecker() {
	exampleQuestion := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	exampleIf := stmt.NewIf(expr.NewIntegerLiteral(10), stmt.NewStmtList([]interfaces.Question{exampleQuestion}, []interfaces.Conditional{}))
	exampleBody := stmt.NewStmtList([]interfaces.Question{}, []interfaces.Conditional{exampleIf})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.NonBooleanConditionError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestDuplicateLabelChecker() {
	firstQuestion := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	secondQuestion := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasMaintLoan"), expr.NewBoolType()))
	exampleBody := stmt.NewStmtList([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfWarningsOfPassedTypeArePresentForForm(exampleForm, errors.DuplicateLabelWarning{}, 1)
}

func (suite *TypeCheckerTestSuite) TestDuplicateVarDeclChecker() {
	firstQuestion := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	secondQuestion := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewIntegerType()))
	exampleBody := stmt.NewStmtList([]interfaces.Question{firstQuestion, secondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.QuestionRedeclaredWithDifferentTypesError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestCyclicReferenceCheckerReferenceToEachOther() {
	questionPointingToSecondQuestion := stmt.NewComputedQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()), expr.NewVarExpr(vari.NewVarID("hasBoughtHouse")))
	questionPointingToFirstQuestion := stmt.NewComputedQuestion(expr.NewStringLiteral("Did you buy a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasBoughtHouse"), expr.NewBoolType()), expr.NewVarExpr(vari.NewVarID("hasSoldHouse")))
	exampleBody := stmt.NewStmtList([]interfaces.Question{questionPointingToFirstQuestion, questionPointingToSecondQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.CyclicDependencyError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestCyclicReferenceCheckerIfConditionRefersToBody() {
	questionExample := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	ifBodyExample := stmt.NewStmtList([]interfaces.Question{questionExample}, []interfaces.Conditional{})
	ifExample := stmt.NewIf(expr.NewVarExpr(vari.NewVarID("hasSoldHouse")), ifBodyExample)
	exampleFormBody := stmt.NewStmtList([]interfaces.Question{}, []interfaces.Conditional{ifExample})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleFormBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.CyclicDependencyError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestQuestionTypeAndComputationTypeMismatch() {
	exampleQuestion := stmt.NewComputedQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()), expr.NewIntegerLiteral(11))
	exampleBody := stmt.NewStmtList([]interfaces.Question{exampleQuestion}, []interfaces.Conditional{})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.DeclaratedTypeAndActualTypeDeviateError{}, 1)
}

func (suite *TypeCheckerTestSuite) TestThatCorrectFormYieldsNoErrorsOrWarnings() {
	firstQuestion := stmt.NewInputQuestion(expr.NewStringLiteral("Did you sell a house in 2010?"), vari.NewVarDecl(vari.NewVarID("hasSoldHouse"), expr.NewBoolType()))
	firstQuestionBody := stmt.NewInputQuestion(expr.NewStringLiteral("What was the selling price?"), vari.NewVarDecl(vari.NewVarID("sellingPrice"), expr.NewIntegerType()))
	ifBody := stmt.NewStmtList([]interfaces.Question{firstQuestionBody}, []interfaces.Conditional{})
	elseBody := stmt.NewStmtList([]interfaces.Question{firstQuestionBody}, []interfaces.Conditional{})
	ifExample := stmt.NewIfElse(expr.NewBoolLiteral(true), ifBody, elseBody)
	exampleBody := stmt.NewStmtList([]interfaces.Question{firstQuestion}, []interfaces.Conditional{ifExample})
	exampleForm := stmt.NewForm(vari.NewVarID("TestForm"), exampleBody)

	suite.testThatNumberOfErrorsOfPassedTypeArePresentForForm(exampleForm, errors.TypeCheckError{}, 0)
}
