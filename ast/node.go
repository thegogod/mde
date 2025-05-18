package ast

import (
	"github.com/thegogod/mde/reflect"
	"github.com/thegogod/mde/tokens"
)

type Node interface {
	Kind() NodeKind
	Start() tokens.Position
	End() tokens.Position
}

type Expr interface {
	Node

	Type() (reflect.Type, error)
	Eval() (*reflect.Value, error)
}

type Stmt interface {
	Node

	Exec() (*reflect.Value, error)
}
