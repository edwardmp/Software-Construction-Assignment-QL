package vari

type VarId struct {
	Ident string
	Var
}

func NewVarId(ident string) VarId {
	return VarId{ident, NewVar()}
}

func (this VarId) GetIdent() string {
	return this.Ident
}

func (this VarId) String() string {
	return this.Ident
}
