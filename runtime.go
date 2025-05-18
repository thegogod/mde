package mde

import (
	"go/ast"

	"github.com/thegogod/mde/reflect"
)

type Runtime interface {
	Run(node ast.Node) (*reflect.Value, error)
	Eval(expr ast.Expr) (*reflect.Value, error)
	Exec(stmt ast.Stmt) (*reflect.Value, error)
}
