package expr

import (
	"fmt"
	"ql/interfaces"
)

/* binary expressions */

func (this Add) String() string {
	return fmt.Sprintf("%s + %s", this.LHS().String(), this.RHS().String())
}

func (this And) String() string {
	return fmt.Sprintf("%s && %s", this.LHS().String(), this.RHS().String())
}

func (this Div) String() string {
	return fmt.Sprintf("%s / %s", this.LHS().String(), this.RHS().String())
}

func (this Eq) String() string {
	return fmt.Sprintf("%s == %s", this.LHS().String(), this.RHS().String())
}

func (this GEq) String() string {
	return fmt.Sprintf("%s >= %s", this.LHS().String(), this.RHS().String())
}

func (this GT) String() string {
	return fmt.Sprintf("%s > %s", this.LHS().String(), this.RHS().String())
}

func (this LEq) String() string {
	return fmt.Sprintf("%s <= %s", this.LHS().String(), this.RHS().String())
}

func (this LT) String() string {
	return fmt.Sprintf("%s < %s", this.LHS().String(), this.RHS().String())
}

func (this Mul) String() string {
	return fmt.Sprintf("%s * %s", this.LHS().String(), this.RHS().String())
}

func (this NEq) String() string {
	return fmt.Sprintf("%s != %s", this.LHS().String(), this.RHS().String())
}

func (this Or) String() string {
	return fmt.Sprintf("%s || %s", this.LHS().String(), this.RHS().String())
}

func (this Sub) String() string {
	return fmt.Sprintf("%s - %s", this.LHS().String(), this.RHS().String())
}

/* literals */

func (this BoolLiteral) String() string {
	return fmt.Sprintf("%t", this.Value())
}

func (this IntegerLiteral) String() string {
	return fmt.Sprintf("%d", this.Value())
}

func (this StringLiteral) String() string {
	return this.Value().(interfaces.StringValue).PrimitiveValueString()
}

/* unary expressions */

func (this Not) String() string {
	return fmt.Sprintf("!%s", this.Value())
}

func (this Pos) String() string {
	return fmt.Sprintf("%s", this.Value())
}

func (this Neg) String() string {
	return fmt.Sprintf("-%s", this.Value())
}

func (this VarExpr) String() string {
	return fmt.Sprintf("%s", this.VarIdentifier())
}

/* value types */

func (this ValueType) String() string {
	return this.typeString
}

/* expressions */

// String is only called on Expr if the struct embedding it did not implement String, which is erroneous
func (this Expr) String() string {
	panic("Expr String method not overridden")
	return ""
}
