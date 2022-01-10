package gui

import (
	"ql/interfaces"
)

type GUIComputedQuestion struct {
	*GUIQuestion
	Expr  interfaces.Expr
	VarID interfaces.VarID
}

// newGUIComputedQuestion is a constructr method returning a new GUIComputedQuestion
func newGUIComputedQuestion(question interfaces.Question, expr interfaces.Expr, varID interfaces.VarID) *GUIComputedQuestion {
	guiQuestion := createDisabledGUIQuestion(question)
	return &GUIComputedQuestion{GUIQuestion: guiQuestion, Expr: expr, VarID: varID}
}
