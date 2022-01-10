package expr

import "ql/interfaces"

/* binary expressions */

func (this Add) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitAdd(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this And) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitAnd(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this Div) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitDiv(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this Eq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitEq(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this GEq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitEq(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this GT) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitGT(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this LEq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitLEq(this, context)

	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this LT) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitLT(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this Mul) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitMul(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this NEq) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitNEq(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this Or) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitOr(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

func (this Sub) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitSub(this, context)
	acceptLHSAndRHS(this, visitor, context)

	return returnValue
}

/* literals */

func (this BoolLiteral) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitBoolLiteral(this, context)

	return returnValue
}

func (this IntegerLiteral) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitIntegerLiteral(this, context)

	return returnValue
}

func (this StringLiteral) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitStringLiteral(this, context)

	return returnValue
}

/* value types */

func (this IntegerType) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitIntegerType(this, context)

	return returnValue
}

func (this BoolType) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitBoolType(this, context)

	return returnValue
}

func (this StringType) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	returnValue := visitor.VisitStringType(this, context)

	return returnValue
}

/* unary expressions */

func (this Not) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitNot(this, context)
	this.Value().Accept(visitor, context)

	return returnValue
}

func (this Pos) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitPos(this, context)
	this.Value().Accept(visitor, context)

	return returnValue
}

func (this VarExpr) Accept(visitor interfaces.Visitor, context interface{}) interface{} {

	returnValue := visitor.VisitVarExpr(this, context)
	this.VarIdentifier().Accept(visitor, context)

	return returnValue
}

// acceptLHSAndRHS provides a convenient way to run Accept on both left-hand and right-hand sides of binary operands
func acceptLHSAndRHS(binaryExpr interfaces.BinaryOperatorExpr, visitor interfaces.Visitor, context interface{}) {
	binaryExpr.LHS().Accept(visitor, context)
	binaryExpr.RHS().Accept(visitor, context)
}

/* expressions */

// Accept is only called on Expr if the struct embedding it did not implement Accept, which is erroneous
func (this Expr) Accept(visitor interfaces.Visitor, context interface{}) interface{} {
	panic("Expr Accept method not overridden")

	return nil
}
