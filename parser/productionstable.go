
package parser

import (
    "ql/ast"
    "ql/token"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Form	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Form : "form" VarID Block	<< ast.NewFormNode(X[1], X[2], X[0].(*token.Token).Pos) >>`,
		Id: "Form",
		NTType: 1,
		Index: 1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewFormNode(X[1], X[2], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Type : "integer"	<< ast.NewIntegerTypeNode(X[0]) >>`,
		Id: "Type",
		NTType: 2,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntegerTypeNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "boolean"	<< ast.NewBoolTypeNode(X[0]) >>`,
		Id: "Type",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolTypeNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Type : "string"	<< ast.NewStringTypeNode(X[0]) >>`,
		Id: "Type",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringTypeNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : integer_lit	<< ast.NewIntegerLiteralNode(X[0]) >>`,
		Id: "Literal",
		NTType: 3,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIntegerLiteralNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Literal : StringLiteraleral	<<  >>`,
		Id: "Literal",
		NTType: 3,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Literal : "true"	<< ast.NewBoolLiteralNode(ast.TRUE, X[0].(*token.Token).Pos) >>`,
		Id: "Literal",
		NTType: 3,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLiteralNode(ast.TRUE, X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Literal : "false"	<< ast.NewBoolLiteralNode(ast.FALSE, X[0].(*token.Token).Pos) >>`,
		Id: "Literal",
		NTType: 3,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewBoolLiteralNode(ast.FALSE, X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `StringLiteraleral : str_lit	<< ast.NewStringLiteralNode(X[0]) >>`,
		Id: "StringLiteraleral",
		NTType: 4,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStringLiteralNode(X[0])
		},
	},
	ProdTabEntry{
		String: `VarDecl : VarID col Type	<< ast.NewVarDeclNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "VarDecl",
		NTType: 5,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarDeclNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `VarID : ident	<< ast.NewVarIDNode(X[0]) >>`,
		Id: "VarID",
		NTType: 6,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarIDNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : addop Expr	<< ast.NewPosNode(X[1], X[0].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPosNode(X[1], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : subop Expr	<< ast.NewNegNode(X[1], X[0].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNegNode(X[1], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : notop Expr	<< ast.NewNotNode(X[1], X[0].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 14,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNotNode(X[1], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr mulop Expr	<< ast.NewMulNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 15,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewMulNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr divop Expr	<< ast.NewDivNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewDivNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr addop Expr	<< ast.NewAddNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAddNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr subop Expr	<< ast.NewSubNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 18,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr eqop Expr	<< ast.NewEqNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEqNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr neqop Expr	<< ast.NewNEqNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNEqNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr gtop Expr	<< ast.NewGTNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGTNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr ltop Expr	<< ast.NewLTNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLTNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr geqop Expr	<< ast.NewGEqNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGEqNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr leqop Expr	<< ast.NewLEqNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewLEqNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr andop Expr	<< ast.NewAndNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAndNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : Expr orop Expr	<< ast.NewOrNode(X[0], X[2], X[1].(*token.Token).Pos) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewOrNode(X[0], X[2], X[1].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Expr : lpar Expr rpar	<< X[1], nil >>`,
		Id: "Expr",
		NTType: 7,
		Index: 27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Expr : VarID	<< ast.NewVarExprNode(X[0]) >>`,
		Id: "Expr",
		NTType: 7,
		Index: 28,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewVarExprNode(X[0])
		},
	},
	ProdTabEntry{
		String: `Expr : Literal	<<  >>`,
		Id: "Expr",
		NTType: 7,
		Index: 29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" lpar Expr rpar Block	<< ast.NewIfNode(X[2], X[4], X[0].(*token.Token).Pos) >>`,
		Id: "IfStmt",
		NTType: 8,
		Index: 30,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfNode(X[2], X[4], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `IfStmt : "if" lpar Expr rpar Block "else" Block	<< ast.NewIfElseNode(X[2], X[4], X[6], X[0].(*token.Token).Pos) >>`,
		Id: "IfStmt",
		NTType: 8,
		Index: 31,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewIfElseNode(X[2], X[4], X[6], X[0].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Question : StringLiteraleral VarDecl	<< ast.NewInputQuestionNode(X[0], X[1]) >>`,
		Id: "Question",
		NTType: 9,
		Index: 32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewInputQuestionNode(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Question : StringLiteraleral VarDecl assign Expr	<< ast.NewComputedQuestionNode(X[0], X[1], X[3], X[2].(*token.Token).Pos) >>`,
		Id: "Question",
		NTType: 9,
		Index: 33,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewComputedQuestionNode(X[0], X[1], X[3], X[2].(*token.Token).Pos)
		},
	},
	ProdTabEntry{
		String: `Stmt : Question	<<  >>`,
		Id: "Stmt",
		NTType: 10,
		Index: 34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : IfStmt	<<  >>`,
		Id: "Stmt",
		NTType: 10,
		Index: 35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `StmtList : Stmt	<< ast.NewStmtListNode(X[0]) >>`,
		Id: "StmtList",
		NTType: 11,
		Index: 36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtListNode(X[0])
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList Stmt	<< ast.AppendStmt(X[0], X[1]) >>`,
		Id: "StmtList",
		NTType: 11,
		Index: 37,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmt(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Block : lbrace StmtList rbrace	<< X[1], nil >>`,
		Id: "Block",
		NTType: 12,
		Index: 38,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Block : lbrace rbrace	<< ast.NewEmptyStmtListNode(X[0].(*token.Token).Pos) >>`,
		Id: "Block",
		NTType: 12,
		Index: 39,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEmptyStmtListNode(X[0].(*token.Token).Pos)
		},
	},
	
}
