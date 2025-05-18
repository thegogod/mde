package ast

import (
	"github.com/thegogod/mde/reflect"
)

type UnaryExpr struct {
	Op    *Token
	Right Expr
}

func NewUnaryExpr(op *Token, right Expr) *UnaryExpr {
	return &UnaryExpr{
		Op:    op,
		Right: right,
	}
}

func (self *UnaryExpr) GetType() (reflect.Type, error) {
	return self.Right.GetType()
}

func (self *UnaryExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitUnaryExpr(self)
}
