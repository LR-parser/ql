package vari

type VarID struct {
	identifier string
	Vari
}

func NewVarID(ident string) VarID {
	return VarID{ident, NewVari()}
}

func (this VarID) Identifier() string {
	return this.identifier
}

func (this VarID) String() string {
	return this.identifier
}
