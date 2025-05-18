package mde

import "github.com/thegogod/mde/ast"

type Parser interface {
	Parse() (ast.Stmt, error)
}
