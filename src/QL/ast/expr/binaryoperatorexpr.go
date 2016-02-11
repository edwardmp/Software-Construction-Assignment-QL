package expr

type BinaryOperatorExpr interface {
	getLhs() interface{}
	getRhs() interface{}
}
