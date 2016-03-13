package vari

type VarId struct {
	Ident string
	Var
}

func NewVarId(ident string, sourceInfo interface{}) VarId {
	return VarId{ident, NewVar(sourceInfo)}
}

func NewVarIdNoSourceInfo(ident string) VarId {
	return NewVarId(ident, nil)
}

func (this VarId) GetIdent() string {
	return this.Ident
}

func (this VarId) String() string {
	return this.Ident
}
