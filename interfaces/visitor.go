package interfaces

type Visitor interface {
	/* Expr */
	VisitAdd(Add, interface{}) interface{}
	VisitAnd(And, interface{}) interface{}
	VisitDiv(Div, interface{}) interface{}
	VisitEq(Eq, interface{}) interface{}
	VisitGEq(GEq, interface{}) interface{}
	VisitGT(GT, interface{}) interface{}
	VisitLEq(LEq, interface{}) interface{}
	VisitLT(LT, interface{}) interface{}
	VisitMul(Mul, interface{}) interface{}
	VisitNEq(NEq, interface{}) interface{}
	VisitOr(Or, interface{}) interface{}
	VisitSub(Sub, interface{}) interface{}
	VisitBoolLiteral(BoolLiteral, interface{}) interface{}
	VisitIntegerLiteral(IntegerLiteral, interface{}) interface{}
	VisitStringLiteral(StringLiteral, interface{}) interface{}
	VisitNeg(Neg, interface{}) interface{}
	VisitNot(Not, interface{}) interface{}
	VisitPos(Pos, interface{}) interface{}
	VisitVarExpr(VarExpr, interface{}) interface{}
	/* Vari */
	VisitVarDecl(VarDecl, interface{}) interface{}
	VisitVarID(VarID, interface{}) interface{}
	VisitIntegerType(IntegerType, interface{}) interface{}
	VisitBoolType(BoolType, interface{}) interface{}
	VisitStringType(StringType, interface{}) interface{}
	/* Stmt */
	VisitForm(Form, interface{}) interface{}
	VisitComputedQuestion(ComputedQuestion, interface{}) interface{}
	VisitInputQuestion(InputQuestion, interface{}) interface{}
	VisitIf(If, interface{}) interface{}
	VisitIfElse(IfElse, interface{}) interface{}
	VisitStmtList(StmtList, interface{}) interface{}
}
