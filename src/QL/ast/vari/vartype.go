package vari

type VarTypeId int

const (
	BOOLEAN VarTypeId = 1 + iota
	STRING
	INT
	MONEY
	DATE
)
