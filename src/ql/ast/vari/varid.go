package vari

type VarId struct {
	Ident string
}

func (va VarId) String() string {
	return va.Ident
}
