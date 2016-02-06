package expr

type Neg struct {
	Value int
}

func (neg Neg) Eval() interface{} {
	return -neg.Value
}
