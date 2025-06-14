package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type Expression interface {
	Type() reflect.Type
	Validate() error
	Evaluate(scope core.Scope) (reflect.Value, error)
}
