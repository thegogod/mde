package html

import "github.com/thegogod/mde/core"

type Node interface {
	core.Node

	GetById(id string) Node
	GetByClass(classes ...string) []Node
}
