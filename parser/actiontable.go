
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			shift(4),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			shift(6),		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			reduce(11),		/* lbrace, reduce: VarID */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Form */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			shift(8),		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(14),		/* rbrace */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			shift(17),		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			reduce(9),		/* ident, reduce: StringLiteraleral */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(35),		/* str_lit, reduce: Stmt */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(35),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(35),		/* rbrace, reduce: Stmt */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(18),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(34),		/* str_lit, reduce: Stmt */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(34),		/* if, reduce: Stmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(34),		/* rbrace, reduce: Stmt */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(36),		/* str_lit, reduce: StmtList */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(36),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(36),		/* rbrace, reduce: StmtList */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			shift(8),		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(20),		/* rbrace */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(39),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			shift(21),		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(32),		/* str_lit, reduce: Question */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(32),		/* if, reduce: Question */
			nil,		/* else */
			shift(22),		/* assign */
			nil,		/* lbrace */
			reduce(32),		/* rbrace, reduce: Question */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			reduce(11),		/* col, reduce: VarID */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(37),		/* str_lit, reduce: StmtList */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(37),		/* if, reduce: StmtList */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(37),		/* rbrace, reduce: StmtList */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(38),		/* $, reduce: Block */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			shift(37),		/* integer */
			shift(38),		/* boolean */
			shift(39),		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(28),		/* addop, reduce: Expr */
			reduce(28),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(28),		/* mulop, reduce: Expr */
			reduce(28),		/* divop, reduce: Expr */
			reduce(28),		/* eqop, reduce: Expr */
			reduce(28),		/* neqop, reduce: Expr */
			reduce(28),		/* gtop, reduce: Expr */
			reduce(28),		/* ltop, reduce: Expr */
			reduce(28),		/* geqop, reduce: Expr */
			reduce(28),		/* leqop, reduce: Expr */
			reduce(28),		/* andop, reduce: Expr */
			reduce(28),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(28),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(29),		/* addop, reduce: Expr */
			reduce(29),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(29),		/* mulop, reduce: Expr */
			reduce(29),		/* divop, reduce: Expr */
			reduce(29),		/* eqop, reduce: Expr */
			reduce(29),		/* neqop, reduce: Expr */
			reduce(29),		/* gtop, reduce: Expr */
			reduce(29),		/* ltop, reduce: Expr */
			reduce(29),		/* geqop, reduce: Expr */
			reduce(29),		/* leqop, reduce: Expr */
			reduce(29),		/* andop, reduce: Expr */
			reduce(29),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(29),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(5),		/* addop, reduce: Literal */
			reduce(5),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(5),		/* mulop, reduce: Literal */
			reduce(5),		/* divop, reduce: Literal */
			reduce(5),		/* eqop, reduce: Literal */
			reduce(5),		/* neqop, reduce: Literal */
			reduce(5),		/* gtop, reduce: Literal */
			reduce(5),		/* ltop, reduce: Literal */
			reduce(5),		/* geqop, reduce: Literal */
			reduce(5),		/* leqop, reduce: Literal */
			reduce(5),		/* andop, reduce: Literal */
			reduce(5),		/* orop, reduce: Literal */
			nil,		/* lpar */
			reduce(5),		/* rpar, reduce: Literal */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(6),		/* addop, reduce: Literal */
			reduce(6),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(6),		/* mulop, reduce: Literal */
			reduce(6),		/* divop, reduce: Literal */
			reduce(6),		/* eqop, reduce: Literal */
			reduce(6),		/* neqop, reduce: Literal */
			reduce(6),		/* gtop, reduce: Literal */
			reduce(6),		/* ltop, reduce: Literal */
			reduce(6),		/* geqop, reduce: Literal */
			reduce(6),		/* leqop, reduce: Literal */
			reduce(6),		/* andop, reduce: Literal */
			reduce(6),		/* orop, reduce: Literal */
			nil,		/* lpar */
			reduce(6),		/* rpar, reduce: Literal */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(7),		/* addop, reduce: Literal */
			reduce(7),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(7),		/* mulop, reduce: Literal */
			reduce(7),		/* divop, reduce: Literal */
			reduce(7),		/* eqop, reduce: Literal */
			reduce(7),		/* neqop, reduce: Literal */
			reduce(7),		/* gtop, reduce: Literal */
			reduce(7),		/* ltop, reduce: Literal */
			reduce(7),		/* geqop, reduce: Literal */
			reduce(7),		/* leqop, reduce: Literal */
			reduce(7),		/* andop, reduce: Literal */
			reduce(7),		/* orop, reduce: Literal */
			nil,		/* lpar */
			reduce(7),		/* rpar, reduce: Literal */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(8),		/* addop, reduce: Literal */
			reduce(8),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(8),		/* mulop, reduce: Literal */
			reduce(8),		/* divop, reduce: Literal */
			reduce(8),		/* eqop, reduce: Literal */
			reduce(8),		/* neqop, reduce: Literal */
			reduce(8),		/* gtop, reduce: Literal */
			reduce(8),		/* ltop, reduce: Literal */
			reduce(8),		/* geqop, reduce: Literal */
			reduce(8),		/* leqop, reduce: Literal */
			reduce(8),		/* andop, reduce: Literal */
			reduce(8),		/* orop, reduce: Literal */
			nil,		/* lpar */
			reduce(8),		/* rpar, reduce: Literal */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: StringLiteraleral */
			reduce(9),		/* subop, reduce: StringLiteraleral */
			nil,		/* notop */
			reduce(9),		/* mulop, reduce: StringLiteraleral */
			reduce(9),		/* divop, reduce: StringLiteraleral */
			reduce(9),		/* eqop, reduce: StringLiteraleral */
			reduce(9),		/* neqop, reduce: StringLiteraleral */
			reduce(9),		/* gtop, reduce: StringLiteraleral */
			reduce(9),		/* ltop, reduce: StringLiteraleral */
			reduce(9),		/* geqop, reduce: StringLiteraleral */
			reduce(9),		/* leqop, reduce: StringLiteraleral */
			reduce(9),		/* andop, reduce: StringLiteraleral */
			reduce(9),		/* orop, reduce: StringLiteraleral */
			nil,		/* lpar */
			reduce(9),		/* rpar, reduce: StringLiteraleral */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: VarID */
			reduce(11),		/* subop, reduce: VarID */
			nil,		/* notop */
			reduce(11),		/* mulop, reduce: VarID */
			reduce(11),		/* divop, reduce: VarID */
			reduce(11),		/* eqop, reduce: VarID */
			reduce(11),		/* neqop, reduce: VarID */
			reduce(11),		/* gtop, reduce: VarID */
			reduce(11),		/* ltop, reduce: VarID */
			reduce(11),		/* geqop, reduce: VarID */
			reduce(11),		/* leqop, reduce: VarID */
			reduce(11),		/* andop, reduce: VarID */
			reduce(11),		/* orop, reduce: VarID */
			nil,		/* lpar */
			reduce(11),		/* rpar, reduce: VarID */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			shift(65),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(10),		/* str_lit, reduce: VarDecl */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(10),		/* if, reduce: VarDecl */
			nil,		/* else */
			reduce(10),		/* assign, reduce: VarDecl */
			nil,		/* lbrace */
			reduce(10),		/* rbrace, reduce: VarDecl */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(2),		/* str_lit, reduce: Type */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(2),		/* if, reduce: Type */
			nil,		/* else */
			reduce(2),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(2),		/* rbrace, reduce: Type */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(3),		/* str_lit, reduce: Type */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(3),		/* if, reduce: Type */
			nil,		/* else */
			reduce(3),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(3),		/* rbrace, reduce: Type */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(4),		/* str_lit, reduce: Type */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(4),		/* if, reduce: Type */
			nil,		/* else */
			reduce(4),		/* assign, reduce: Type */
			nil,		/* lbrace */
			reduce(4),		/* rbrace, reduce: Type */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(28),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			reduce(28),		/* addop, reduce: Expr */
			reduce(28),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(28),		/* mulop, reduce: Expr */
			reduce(28),		/* divop, reduce: Expr */
			reduce(28),		/* eqop, reduce: Expr */
			reduce(28),		/* neqop, reduce: Expr */
			reduce(28),		/* gtop, reduce: Expr */
			reduce(28),		/* ltop, reduce: Expr */
			reduce(28),		/* geqop, reduce: Expr */
			reduce(28),		/* leqop, reduce: Expr */
			reduce(28),		/* andop, reduce: Expr */
			reduce(28),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(28),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(28),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(29),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			reduce(29),		/* addop, reduce: Expr */
			reduce(29),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(29),		/* mulop, reduce: Expr */
			reduce(29),		/* divop, reduce: Expr */
			reduce(29),		/* eqop, reduce: Expr */
			reduce(29),		/* neqop, reduce: Expr */
			reduce(29),		/* gtop, reduce: Expr */
			reduce(29),		/* ltop, reduce: Expr */
			reduce(29),		/* geqop, reduce: Expr */
			reduce(29),		/* leqop, reduce: Expr */
			reduce(29),		/* andop, reduce: Expr */
			reduce(29),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(29),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(29),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(5),		/* str_lit, reduce: Literal */
			nil,		/* col */
			nil,		/* ident */
			reduce(5),		/* addop, reduce: Literal */
			reduce(5),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(5),		/* mulop, reduce: Literal */
			reduce(5),		/* divop, reduce: Literal */
			reduce(5),		/* eqop, reduce: Literal */
			reduce(5),		/* neqop, reduce: Literal */
			reduce(5),		/* gtop, reduce: Literal */
			reduce(5),		/* ltop, reduce: Literal */
			reduce(5),		/* geqop, reduce: Literal */
			reduce(5),		/* leqop, reduce: Literal */
			reduce(5),		/* andop, reduce: Literal */
			reduce(5),		/* orop, reduce: Literal */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(5),		/* if, reduce: Literal */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(5),		/* rbrace, reduce: Literal */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(6),		/* str_lit, reduce: Literal */
			nil,		/* col */
			nil,		/* ident */
			reduce(6),		/* addop, reduce: Literal */
			reduce(6),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(6),		/* mulop, reduce: Literal */
			reduce(6),		/* divop, reduce: Literal */
			reduce(6),		/* eqop, reduce: Literal */
			reduce(6),		/* neqop, reduce: Literal */
			reduce(6),		/* gtop, reduce: Literal */
			reduce(6),		/* ltop, reduce: Literal */
			reduce(6),		/* geqop, reduce: Literal */
			reduce(6),		/* leqop, reduce: Literal */
			reduce(6),		/* andop, reduce: Literal */
			reduce(6),		/* orop, reduce: Literal */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(6),		/* if, reduce: Literal */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(6),		/* rbrace, reduce: Literal */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(7),		/* str_lit, reduce: Literal */
			nil,		/* col */
			nil,		/* ident */
			reduce(7),		/* addop, reduce: Literal */
			reduce(7),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(7),		/* mulop, reduce: Literal */
			reduce(7),		/* divop, reduce: Literal */
			reduce(7),		/* eqop, reduce: Literal */
			reduce(7),		/* neqop, reduce: Literal */
			reduce(7),		/* gtop, reduce: Literal */
			reduce(7),		/* ltop, reduce: Literal */
			reduce(7),		/* geqop, reduce: Literal */
			reduce(7),		/* leqop, reduce: Literal */
			reduce(7),		/* andop, reduce: Literal */
			reduce(7),		/* orop, reduce: Literal */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(7),		/* if, reduce: Literal */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(7),		/* rbrace, reduce: Literal */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(8),		/* str_lit, reduce: Literal */
			nil,		/* col */
			nil,		/* ident */
			reduce(8),		/* addop, reduce: Literal */
			reduce(8),		/* subop, reduce: Literal */
			nil,		/* notop */
			reduce(8),		/* mulop, reduce: Literal */
			reduce(8),		/* divop, reduce: Literal */
			reduce(8),		/* eqop, reduce: Literal */
			reduce(8),		/* neqop, reduce: Literal */
			reduce(8),		/* gtop, reduce: Literal */
			reduce(8),		/* ltop, reduce: Literal */
			reduce(8),		/* geqop, reduce: Literal */
			reduce(8),		/* leqop, reduce: Literal */
			reduce(8),		/* andop, reduce: Literal */
			reduce(8),		/* orop, reduce: Literal */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(8),		/* if, reduce: Literal */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(8),		/* rbrace, reduce: Literal */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(9),		/* str_lit, reduce: StringLiteraleral */
			nil,		/* col */
			nil,		/* ident */
			reduce(9),		/* addop, reduce: StringLiteraleral */
			reduce(9),		/* subop, reduce: StringLiteraleral */
			nil,		/* notop */
			reduce(9),		/* mulop, reduce: StringLiteraleral */
			reduce(9),		/* divop, reduce: StringLiteraleral */
			reduce(9),		/* eqop, reduce: StringLiteraleral */
			reduce(9),		/* neqop, reduce: StringLiteraleral */
			reduce(9),		/* gtop, reduce: StringLiteraleral */
			reduce(9),		/* ltop, reduce: StringLiteraleral */
			reduce(9),		/* geqop, reduce: StringLiteraleral */
			reduce(9),		/* leqop, reduce: StringLiteraleral */
			reduce(9),		/* andop, reduce: StringLiteraleral */
			reduce(9),		/* orop, reduce: StringLiteraleral */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(9),		/* if, reduce: StringLiteraleral */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(9),		/* rbrace, reduce: StringLiteraleral */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(11),		/* str_lit, reduce: VarID */
			nil,		/* col */
			nil,		/* ident */
			reduce(11),		/* addop, reduce: VarID */
			reduce(11),		/* subop, reduce: VarID */
			nil,		/* notop */
			reduce(11),		/* mulop, reduce: VarID */
			reduce(11),		/* divop, reduce: VarID */
			reduce(11),		/* eqop, reduce: VarID */
			reduce(11),		/* neqop, reduce: VarID */
			reduce(11),		/* gtop, reduce: VarID */
			reduce(11),		/* ltop, reduce: VarID */
			reduce(11),		/* geqop, reduce: VarID */
			reduce(11),		/* leqop, reduce: VarID */
			reduce(11),		/* andop, reduce: VarID */
			reduce(11),		/* orop, reduce: VarID */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(11),		/* if, reduce: VarID */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(11),		/* rbrace, reduce: VarID */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(33),		/* str_lit, reduce: Question */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(33),		/* if, reduce: Question */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(33),		/* rbrace, reduce: Question */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S61
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S62
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S63
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S64
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(25),		/* integer_lit */
			shift(27),		/* true */
			shift(28),		/* false */
			shift(29),		/* str_lit */
			nil,		/* col */
			shift(30),		/* ident */
			shift(32),		/* addop */
			shift(33),		/* subop */
			shift(34),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(35),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S65
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			shift(99),		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S66
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(12),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S67
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(13),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S68
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(14),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S69
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			shift(100),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S70
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S71
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S72
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S73
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S74
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S75
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S76
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S77
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S78
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S79
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S80
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S81
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			shift(42),		/* integer_lit */
			shift(44),		/* true */
			shift(45),		/* false */
			shift(46),		/* str_lit */
			nil,		/* col */
			shift(47),		/* ident */
			shift(49),		/* addop */
			shift(50),		/* subop */
			shift(51),		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			shift(52),		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S82
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(12),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(12),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(12),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S83
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(13),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(13),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(13),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S84
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(14),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(14),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(14),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S85
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			shift(113),		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S86
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(17),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S87
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(18),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S88
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(15),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S89
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(16),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S90
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(19),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S91
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(20),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S92
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(21),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S93
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(22),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S94
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(23),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S95
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(24),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S96
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(25),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S97
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			shift(53),		/* addop */
			shift(54),		/* subop */
			nil,		/* notop */
			shift(55),		/* mulop */
			shift(56),		/* divop */
			shift(57),		/* eqop */
			shift(58),		/* neqop */
			shift(59),		/* gtop */
			shift(60),		/* ltop */
			shift(61),		/* geqop */
			shift(62),		/* leqop */
			shift(63),		/* andop */
			shift(64),		/* orop */
			nil,		/* lpar */
			reduce(26),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S98
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(30),		/* str_lit, reduce: IfStmt */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(30),		/* if, reduce: IfStmt */
			shift(114),		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(30),		/* rbrace, reduce: IfStmt */
			
		},

	},
	actionRow{ // S99
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			shift(8),		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(116),		/* rbrace */
			
		},

	},
	actionRow{ // S100
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			reduce(27),		/* addop, reduce: Expr */
			reduce(27),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(27),		/* mulop, reduce: Expr */
			reduce(27),		/* divop, reduce: Expr */
			reduce(27),		/* eqop, reduce: Expr */
			reduce(27),		/* neqop, reduce: Expr */
			reduce(27),		/* gtop, reduce: Expr */
			reduce(27),		/* ltop, reduce: Expr */
			reduce(27),		/* geqop, reduce: Expr */
			reduce(27),		/* leqop, reduce: Expr */
			reduce(27),		/* andop, reduce: Expr */
			reduce(27),		/* orop, reduce: Expr */
			nil,		/* lpar */
			reduce(27),		/* rpar, reduce: Expr */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S101
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(17),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(17),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(17),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S102
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(18),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(18),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(18),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S103
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(15),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(15),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(15),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S104
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(16),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(16),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(16),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S105
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(19),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(19),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(19),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S106
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(20),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(20),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(20),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S107
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(21),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(21),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(21),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S108
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(22),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(22),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(22),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S109
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(23),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(23),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(23),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S110
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(24),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(24),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(24),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S111
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(25),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(25),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(25),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S112
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(26),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			shift(70),		/* addop */
			shift(71),		/* subop */
			nil,		/* notop */
			shift(72),		/* mulop */
			shift(73),		/* divop */
			shift(74),		/* eqop */
			shift(75),		/* neqop */
			shift(76),		/* gtop */
			shift(77),		/* ltop */
			shift(78),		/* geqop */
			shift(79),		/* leqop */
			shift(80),		/* andop */
			shift(81),		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(26),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(26),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S113
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(27),		/* str_lit, reduce: Expr */
			nil,		/* col */
			nil,		/* ident */
			reduce(27),		/* addop, reduce: Expr */
			reduce(27),		/* subop, reduce: Expr */
			nil,		/* notop */
			reduce(27),		/* mulop, reduce: Expr */
			reduce(27),		/* divop, reduce: Expr */
			reduce(27),		/* eqop, reduce: Expr */
			reduce(27),		/* neqop, reduce: Expr */
			reduce(27),		/* gtop, reduce: Expr */
			reduce(27),		/* ltop, reduce: Expr */
			reduce(27),		/* geqop, reduce: Expr */
			reduce(27),		/* leqop, reduce: Expr */
			reduce(27),		/* andop, reduce: Expr */
			reduce(27),		/* orop, reduce: Expr */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(27),		/* if, reduce: Expr */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(27),		/* rbrace, reduce: Expr */
			
		},

	},
	actionRow{ // S114
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			nil,		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			nil,		/* if */
			nil,		/* else */
			nil,		/* assign */
			shift(118),		/* lbrace */
			nil,		/* rbrace */
			
		},

	},
	actionRow{ // S115
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			shift(8),		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(119),		/* rbrace */
			
		},

	},
	actionRow{ // S116
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(39),		/* str_lit, reduce: Block */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(39),		/* if, reduce: Block */
			reduce(39),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(39),		/* rbrace, reduce: Block */
			
		},

	},
	actionRow{ // S117
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(31),		/* str_lit, reduce: IfStmt */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(31),		/* if, reduce: IfStmt */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(31),		/* rbrace, reduce: IfStmt */
			
		},

	},
	actionRow{ // S118
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			shift(8),		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(121),		/* rbrace */
			
		},

	},
	actionRow{ // S119
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(38),		/* str_lit, reduce: Block */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(38),		/* if, reduce: Block */
			reduce(38),		/* else, reduce: Block */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: Block */
			
		},

	},
	actionRow{ // S120
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			shift(8),		/* str_lit */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			shift(10),		/* if */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			shift(122),		/* rbrace */
			
		},

	},
	actionRow{ // S121
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(39),		/* str_lit, reduce: Block */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(39),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(39),		/* rbrace, reduce: Block */
			
		},

	},
	actionRow{ // S122
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* form */
			nil,		/* integer */
			nil,		/* boolean */
			nil,		/* string */
			nil,		/* integer_lit */
			nil,		/* true */
			nil,		/* false */
			reduce(38),		/* str_lit, reduce: Block */
			nil,		/* col */
			nil,		/* ident */
			nil,		/* addop */
			nil,		/* subop */
			nil,		/* notop */
			nil,		/* mulop */
			nil,		/* divop */
			nil,		/* eqop */
			nil,		/* neqop */
			nil,		/* gtop */
			nil,		/* ltop */
			nil,		/* geqop */
			nil,		/* leqop */
			nil,		/* andop */
			nil,		/* orop */
			nil,		/* lpar */
			nil,		/* rpar */
			reduce(38),		/* if, reduce: Block */
			nil,		/* else */
			nil,		/* assign */
			nil,		/* lbrace */
			reduce(38),		/* rbrace, reduce: Block */
			
		},

	},
	
}

