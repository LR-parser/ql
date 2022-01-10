package stmt

import "ql/interfaces"

type Form struct {
	identifier interfaces.VarID
	content    StmtList
	Stmt
}

func NewForm(identifier interfaces.VarID, content StmtList) Form {
	return Form{identifier: identifier, content: content, Stmt: NewStmt()}
}

func (this Form) Content() interfaces.StmtList {
	return this.content
}

func (this Form) Questions() []interfaces.Question {
	return this.content.Questions()
}

func (this Form) Identifier() interfaces.VarID {
	return this.identifier
}
