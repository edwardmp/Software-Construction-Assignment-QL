package node

import (
	"ql/token"
)

type Node struct {
	SourceInfo token.Pos
}

func NewNode() Node {
	return Node{SourceInfo: token.Pos{}}
}

func (this Node) SetSourceInfo(sourceInfo token.Pos) {
	this.SourceInfo = token.Pos{}
}

func (this Node) GetSourceInfo() token.Pos {
	return this.SourceInfo
}
