package stmt

type Conditional interface {
	EvalCondition() bool
}
