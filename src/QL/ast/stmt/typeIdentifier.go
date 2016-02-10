package stmt

type TypeIdentifier int

const (
	BOOLEAN TypeIdentifier = 1 + iota
	STRING
	INT
	MONEY
	DATE
)
