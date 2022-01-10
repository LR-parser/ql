package interfaces

import "ql/token"

type Node interface {
	SourceInfo() token.Pos
}
