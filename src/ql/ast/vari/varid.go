package vari

type VarId struct {
	identifier string
	Vari
}

func NewVarId(ident string) VarId {
	return VarId{ident, NewVari()}
}

func (this VarId) Identifier() string {
	return this.identifier
}

func (this VarId) String() string {
	return this.identifier
}
