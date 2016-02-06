package expr

type Expr interface {
	Eval() interface{}
}
