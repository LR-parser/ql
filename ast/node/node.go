package node

import "ql/token"

type Node struct {
	SrcInfo token.Pos
}

func NewNode() Node {
	return Node{SrcInfo: token.Pos{}}
}

func (this Node) SetSourceInfo(sourceInfo token.Pos) {
	this.SrcInfo = token.Pos{}
}

func (this Node) SourceInfo() token.Pos {
	return this.SrcInfo
}

func (this Node) String() string {
	return ""
}
