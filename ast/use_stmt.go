package ast

import (
	"github.com/thegogod/mde/reflect"
)

type UseStmt struct {
	Path  []*Token
	Stmts []Stmt
}

func NewUseStmt(path []*Token, stmts []Stmt) *UseStmt {
	return &UseStmt{
		Path:  path,
		Stmts: stmts,
	}
}

func (self *UseStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitUseStmt(self)
}
