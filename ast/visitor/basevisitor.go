package visitor

import "ql/interfaces"

type BaseVisitor struct {
}

func NewBaseVisitor() *BaseVisitor {
	return new(BaseVisitor)
}

func (this *BaseVisitor) VisitAdd(a interfaces.Add, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitAnd(a interfaces.And, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitDiv(d interfaces.Div, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitEq(e interfaces.Eq, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitGEq(g interfaces.GEq, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitGT(g interfaces.GT, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitLEq(l interfaces.LEq, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitLT(l interfaces.LT, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitMul(m interfaces.Mul, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitNEq(n interfaces.NEq, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitOr(o interfaces.Or, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitSub(su interfaces.Sub, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitBoolLiteral(bo interfaces.BoolLiteral, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitIntegerLiteral(i interfaces.IntegerLiteral, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitStringLiteral(st interfaces.StringLiteral, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitNeg(n interfaces.Neg, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitNot(n interfaces.Not, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitPos(p interfaces.Pos, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitVarExpr(va interfaces.VarExpr, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitVarDecl(va interfaces.VarDecl, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitVarID(va interfaces.VarID, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitIntegerType(i interfaces.IntegerType, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitBoolType(bo interfaces.BoolType, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitStringType(st interfaces.StringType, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitForm(f interfaces.Form, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitComputedQuestion(c interfaces.ComputedQuestion, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitInputQuestion(i interfaces.InputQuestion, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitIf(i interfaces.If, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitIfElse(i interfaces.IfElse, context interface{}) interface{} {
	return nil
}

func (this *BaseVisitor) VisitStmtList(st interfaces.StmtList, context interface{}) interface{} {
	return nil
}
