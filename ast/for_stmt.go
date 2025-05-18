package ast

import (
	"github.com/thegogod/mde/reflect"
)

type ForStmt struct {
	Init Stmt
	Cond Expr
	Inc  Expr
	Body Stmt
}

func NewForStmt(init Stmt, cond Expr, inc Expr, body Stmt) *ForStmt {
	return &ForStmt{
		Init: init,
		Cond: cond,
		Inc:  inc,
		Body: body,
	}
}

func (self *ForStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitForStmt(self)
}
