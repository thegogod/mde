package ast

import (
	"github.com/thegogod/mde/reflect"
	"github.com/thegogod/mde/tokens"
)

type Node interface {
	Kind() NodeKind
	Start() tokens.Position
	End() tokens.Position
	Exec() (*reflect.Value, error)
}

type Expr interface {
	Node

	Type() (reflect.Type, error)
}

type Stmt interface {
	Node
}
