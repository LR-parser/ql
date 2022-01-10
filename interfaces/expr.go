package interfaces

type Expr interface {
	ASTNode
	TypeCheck(TypeCheckArgs) ValueType
	String() string
	Eval(VarIDValueSymbols) Value
}

/** value */

type Value interface {
	PrimitiveValue() interface{}
	Eq(value Value) Value
	NEq(value Value) Value
	GEq(value Value) Value
	LEq(value Value) Value
	GT(value Value) Value
	LT(value Value) Value
	Div(value Value) Value
	Mul(value Value) Value
	Add(value Value) Value
	Sub(value Value) Value
	Neg() Value
	Pos() Value
	And(value Value) Value
	Or(value Value) Value
	Not() Value
}

type BoolValue interface {
	PrimitiveValueBool() bool
}

type StringValue interface {
	PrimitiveValueString() string
}

/** binary operator expressions */

type BinaryOperatorExpr interface {
	Expr
	LHS() Expr
	RHS() Expr
}

type Add interface {
	BinaryOperatorExpr
}

type And interface {
	BinaryOperatorExpr
}

type Div interface {
	BinaryOperatorExpr
}

type Eq interface {
	BinaryOperatorExpr
}

type GEq interface {
	BinaryOperatorExpr
}

type GT interface {
	BinaryOperatorExpr
}

type LEq interface {
	BinaryOperatorExpr
}

type LT interface {
	BinaryOperatorExpr
}

type Mul interface {
	BinaryOperatorExpr
}

type NEq interface {
	BinaryOperatorExpr
}

type Or interface {
	BinaryOperatorExpr
}

type Sub interface {
	BinaryOperatorExpr
}

/** unary operator expressions **/

type UnaryOperatorExpr interface {
	Expr
	Value() Expr
}

type Neg interface {
	UnaryOperatorExpr
}

type Not interface {
	UnaryOperatorExpr
}

type Pos interface {
	UnaryOperatorExpr
}

type VarExpr interface {
	VarIdentifier() VarID
}

/** value types **/

type ValueType interface {
	Expr
	DefaultValue() LitExpr
}

type UnknownType interface {
	ValueType
}

type IntegerType interface {
	ValueType
}

type BoolType interface {
	ValueType
}

type StringType interface {
	ValueType
}

/** literals **/

type LitExpr interface {
	Expr
}

type StringLiteral interface {
	LitExpr
	Value() Value
}

type BoolLiteral interface {
	LitExpr
	Value() Value
}

type IntegerLiteral interface {
	LitExpr
	Value() Value
}
