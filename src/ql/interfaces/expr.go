package interfaces

type Expr interface {
	ASTNode
	Eval(s VarIdValueSymbols) interface{}
	TypeCheck(TypeCheckArgs) ValueType
}

type BinaryOperatorExpr interface {
	Expr
	GetLhs() Expr
	GetRhs() Expr
}

type Add interface {
	BinaryOperatorExpr
}

type And interface {
	BinaryOperatorExpr
}

type Div interface {
	BinaryOperatorExpr
}

type Eq interface {
	BinaryOperatorExpr
}

type GEq interface {
	BinaryOperatorExpr
}

type GT interface {
	BinaryOperatorExpr
}

type LEq interface {
	BinaryOperatorExpr
}

type LT interface {
	BinaryOperatorExpr
}

type Mul interface {
	BinaryOperatorExpr
}

type NEq interface {
	BinaryOperatorExpr
}

type Or interface {
	BinaryOperatorExpr
}

type Sub interface {
	BinaryOperatorExpr
}

type LitExpr interface {
	Expr
}

type StrLit interface {
	LitExpr
	GetValue() string
}

type BoolLit interface {
	LitExpr
	GetValue() bool
}

type IntLit interface {
	LitExpr
	GetValue() int
}

type UnaryOperatorExpr interface {
	Expr
	GetValue() Expr
}

type Neg interface {
	UnaryOperatorExpr
}

type Not interface {
	UnaryOperatorExpr
}

type Pos interface {
	UnaryOperatorExpr
}

type VarExpr interface {
	// TODO body?
}

type ValueType interface {
	Expr
	GetDefaultValue() LitExpr
}

type UnknownType interface {
	ValueType
}

type IntType interface {
	ValueType
}

type BoolType interface {
	ValueType
}

type StringType interface {
	ValueType
}
