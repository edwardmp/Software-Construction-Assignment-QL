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

func (v VarId) GetIdent() string {
	return v.Ident
}

func (va VarId) String() string {
	return va.Ident
}
