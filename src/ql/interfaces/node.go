package interfaces

import "ql/token"

type Node interface {
	GetSourceInfo() token.Pos
}
