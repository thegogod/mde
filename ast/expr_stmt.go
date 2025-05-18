package ast

import (
	"github.com/thegogod/mde/reflect"
)

type ExprStmt struct {
	Expr Expr
}

func NewExprStmt(expr Expr) *ExprStmt {
	return &ExprStmt{
		Expr: expr,
	}
}

func (self *ExprStmt) Accept(v StmtVisitor) (*reflect.Value, error) {
	return v.VisitExprStmt(self)
}
