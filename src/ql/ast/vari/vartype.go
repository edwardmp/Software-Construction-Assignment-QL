package vari

type VarType int

const (
	BOOLEAN VarType = 1 + iota
	STRING
	INT
	MONEY
	DATE
)
