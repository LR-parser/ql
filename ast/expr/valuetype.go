package expr

type ValueType struct {
	Expr
	typeString string
}

func NewValueType(typeString string) ValueType {
	return ValueType{Expr: NewExpr(), typeString: typeString}
}
