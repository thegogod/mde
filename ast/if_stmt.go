package ast

import (
	"github.com/thegogod/mde/reflect"
)

type IfStmt struct {
	Cond Expr
	Then Stmt
	Else Stmt
}

func NewIfStmt(cond Expr, then Stmt, _else Stmt) *IfStmt {
	return &IfStmt{
		Cond: cond,
		Then: then,
		Else: _else,
	}
}

func (self *IfStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitIfStmt(self)
}
