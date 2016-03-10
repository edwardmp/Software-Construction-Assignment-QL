package interfaces

type UnaryOperatorExpr interface {
	Expr
	GetValue() string
}
