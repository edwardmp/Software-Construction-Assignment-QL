package expr

type Expr interface {
	Eval(s interface{}) interface{}
}
