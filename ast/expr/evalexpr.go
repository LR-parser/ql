package expr

import "ql/interfaces"

func (this Add) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).Add(this.RHS().Eval(symbols))
}

func (this And) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).And(this.RHS().Eval(symbols))
}

func (this Div) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).Div(this.RHS().Eval(symbols))
}

func (this Eq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).Eq(this.RHS().Eval(symbols))
}

func (this GEq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).GEq(this.RHS().Eval(symbols))
}

func (this GT) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).GT(this.RHS().Eval(symbols))
}

func (this LEq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).LEq(this.RHS().Eval(symbols))
}

func (this LT) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).LT(this.RHS().Eval(symbols))
}

func (this Mul) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).Mul(this.RHS().Eval(symbols))
}

func (this Neg) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value().Eval(symbols).Neg()
}

func (this NEq) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).NEq(this.RHS().Eval(symbols))
}

func (this Not) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value().Eval(symbols).Not()
}

func (this Or) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).Or(this.RHS().Eval(symbols))
}

func (this Pos) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value().Eval(symbols).Pos()
}

func (this IntegerLiteral) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value()
}

func (this StringLiteral) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value()
}

func (this BoolLiteral) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.Value()
}

func (this Sub) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	return this.LHS().Eval(symbols).Sub(this.RHS().Eval(symbols))
}

func (this VarExpr) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	if symbols == nil {
		panic("No symbols passed to Eval VarExpr")
	}

	referencedExpr := symbols.ExprForVarID(this.VarIdentifier())
	if referencedExpr == nil {
		panic("VarExpr refers to nil expression")
	}

	return referencedExpr.Eval(symbols)
}

func (this Expr) Eval(symbols interfaces.VarIDValueSymbols) interfaces.Value {
	panic("Expr Eval method not overridden")

	return nil
}
