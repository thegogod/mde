package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type Statement interface {
	core.Node

	Validate() error
	Evaluate(scope core.Scope) (reflect.Value, error)
}
