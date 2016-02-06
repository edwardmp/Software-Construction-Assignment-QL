package expr

type Pos struct {
	Value int
}

func (pos Pos) Eval() interface{} {
	return +pos.Value
}
