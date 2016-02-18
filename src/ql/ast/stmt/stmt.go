package stmt

type Stmt interface {
	Eval()
	String() string
}
