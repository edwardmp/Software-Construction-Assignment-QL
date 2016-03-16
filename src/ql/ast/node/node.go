package node

import (
	"ql/token"
)

type Node struct {
	SourceInfo token.Pos
}

func NewNode(sourceInfo interface{}) Node {
	if sourceInfo == nil {
		sourceInfo = token.Pos{}
	}

	return Node{}
	//return Node{sourceInfo.(token.Pos)}
}

func (this Node) GetSourceInfo() token.Pos {
	return this.SourceInfo
}
