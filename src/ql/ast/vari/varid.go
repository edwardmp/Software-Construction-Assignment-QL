package vari

type VarId struct {
	Ident string
}

func (v VarId) GetIdent() string {
	return v.Ident
}

func (va VarId) String() string {
	return va.Ident
}
