package vari

type VarId struct {
	identifier string
	Var
}

func NewVarId(ident string) VarId {
	return VarId{ident, NewVar()}
}

func (this VarId) Identifier() string {
	return this.identifier
}

func (this VarId) String() string {
	return this.identifier
}
