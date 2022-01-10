package stmt

import "ql/interfaces"

type InputQuestion struct {
	Question
}

func NewInputQuestion(label interfaces.StringLiteral, varDecl interfaces.VarDecl) InputQuestion {
	return InputQuestion{Question: NewQuestion(label, varDecl)}
}
