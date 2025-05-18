package ast

import (
	"github.com/thegogod/mde/reflect"
)

type AssignExpr struct {
	Name  *Token
	Value Expr
}

func NewAssignExpr(name *Token, value Expr) *AssignExpr {
	return &AssignExpr{
		Name:  name,
		Value: value,
	}
}

func (self *AssignExpr) GetType() (reflect.Type, error) {
	return self.Value.GetType()
}

func (self *AssignExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitAssignExpr(self)
}
