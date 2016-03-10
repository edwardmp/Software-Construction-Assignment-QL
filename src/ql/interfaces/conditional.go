package interfaces

type Conditional interface {
	Node
	EvalCondition() bool
}
