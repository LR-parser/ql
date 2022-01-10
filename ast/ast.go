package ast

import (
	"ql/ast/expr"
	"ql/ast/stmt"
	"ql/ast/vari"
	"ql/interfaces"
	"ql/token"
	"ql/util"
)

const (
	TRUE  = true
	FALSE = false
)

type attrib interface {
}

/** Expressions **/

/* unary operator expressions */

func NewPosNode(value attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewPos(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNegNode(value attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewNeg(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNotNode(value attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewNot(value.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewVarExprNode(identifier attrib) (interfaces.Expr, error) {
	varID := identifier.(vari.VarID)
	expr := expr.NewVarExpr(varID)
	expr.SetSourceInfo(varID.SourceInfo())
	return expr, nil
}

/* binary operator expressions */

func NewMulNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewMul(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewDivNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewDiv(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewAddNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewAdd(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewSubNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewSub(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewEqNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewNEqNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewNEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewGTNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewGT(lhs.(interfaces.Expr).(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewLTNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewLT(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewGEqNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewGEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewLEqNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewLEq(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewAndNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewAnd(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewOrNode(lhs attrib, rhs attrib, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewOr(lhs.(interfaces.Expr), rhs.(interfaces.Expr))
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

/* Literals */

func NewIntegerLiteralNode(litValueToken attrib) (interfaces.Expr, error) {
	sourcePosInfo := litValueToken.(*token.Token).Pos
	value, err := util.IntValue(litValueToken.(*token.Token).Lit)
	expr := expr.NewIntegerLiteral(int(value))
	expr.SetSourceInfo(sourcePosInfo)
	return expr, err
}

func NewBoolLiteralNode(value bool, sourcePosInfo attrib) (interfaces.Expr, error) {
	expr := expr.NewBoolLiteral(value)
	expr.SetSourceInfo(sourcePosInfo.(token.Pos))
	return expr, nil
}

func NewStringLiteralNode(valueToken attrib) (interfaces.Expr, error) {
	sourcePosInfo := valueToken.(*token.Token).Pos
	literalString := util.StringLiteralTokensToString(valueToken.(*token.Token))
	expr := expr.NewStringLiteral(literalString)
	expr.SetSourceInfo(sourcePosInfo)
	return expr, nil
}

/* Value types */

func NewIntegerTypeNode(typeTokenLit attrib) (interfaces.IntegerType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewIntegerType()
	expr.SetSourceInfo(token.Pos)

	return expr, nil
}

func NewBoolTypeNode(typeTokenLit attrib) (interfaces.BoolType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewBoolType()
	expr.SetSourceInfo(token.Pos)
	return expr, nil
}

func NewStringTypeNode(typeTokenLit attrib) (interfaces.StringType, error) {
	token := typeTokenLit.(*token.Token)
	expr := expr.NewStringType()
	expr.SetSourceInfo(token.Pos)
	return expr, nil
}

/** Vari **/

func NewVarDeclNode(ident attrib, typeIdent attrib, sourcePosInfo attrib) (interfaces.VarDecl, error) {
	vari := vari.NewVarDecl(ident.(interfaces.VarID), typeIdent.(interfaces.ValueType))
	vari.SetSourceInfo(sourcePosInfo.(token.Pos))
	return vari, nil
}

func NewVarIDNode(identToken attrib) (vari.VarID, error) {
	sourcePosInfo := identToken.(*token.Token).Pos
	identifierString := string(identToken.(*token.Token).Lit)
	vari := vari.NewVarID(identifierString)
	vari.SetSourceInfo(sourcePosInfo)
	return vari, nil
}

/** Statements **/

func NewFormNode(identifier attrib, body attrib, sourcePosInfo attrib) (interfaces.Form, error) {
	stmt := stmt.NewForm(identifier.(vari.VarID), body.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewInputQuestionNode(label attrib, varDecl attrib) (interfaces.InputQuestion, error) {
	labelStringLiteral := label.(expr.StringLiteral)
	stmt := stmt.NewInputQuestion(labelStringLiteral, varDecl.(vari.VarDecl))
	stmt.SetSourceInfo(labelStringLiteral.SourceInfo())
	return stmt, nil
}

func NewComputedQuestionNode(label attrib, varDecl attrib, computation attrib, sourcePosInfo attrib) (interfaces.ComputedQuestion, error) {
	stmt := stmt.NewComputedQuestion(label.(expr.StringLiteral), varDecl.(vari.VarDecl), computation.(interfaces.Expr))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewStmtListNode(stmtElt attrib) (interfaces.StmtList, error) {
	stmtEltTypeAsserted := stmtElt.(interfaces.Stmt)
	stmt := stmt.NewEmptyStmtList()
	stmt.SetSourceInfo(stmtEltTypeAsserted.SourceInfo())
	return stmt.AddStmt(stmtEltTypeAsserted), nil
}

func NewEmptyStmtListNode(sourcePosInfo attrib) (interfaces.StmtList, error) {
	stmt := stmt.NewEmptyStmtList()
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func AppendStmt(stmtList, stmtElt attrib) (interfaces.StmtList, error) {
	stmt := stmtList.(stmt.StmtList).AddStmt(stmtElt.(interfaces.Stmt))
	return stmt, nil
}

func NewIfNode(cond attrib, body attrib, sourcePosInfo attrib) (interfaces.If, error) {
	stmt := stmt.NewIf(cond.(interfaces.Expr), body.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}

func NewIfElseNode(cond attrib, ifBody attrib, elseBody attrib, sourcePosInfo attrib) (interfaces.IfElse, error) {
	stmt := stmt.NewIfElse(cond.(interfaces.Expr), ifBody.(stmt.StmtList), elseBody.(stmt.StmtList))
	stmt.SetSourceInfo(sourcePosInfo.(token.Pos))
	return stmt, nil
}
