package ast

import (
	"github.com/thegogod/mde/reflect"
)

type CallExpr struct {
	Callee Expr
	Paren  *Token
	Args   []Expr
}

func NewCallExpr(callee Expr, paren *Token, args []Expr) *CallExpr {
	return &CallExpr{
		Callee: callee,
		Paren:  paren,
		Args:   args,
	}
}

func (self *CallExpr) GetType() (reflect.Type, error) {
	t, err := self.Callee.GetType()

	if err != nil {
		return nil, err
	}

	if callable, ok := t.(reflect.CallableType); ok {
		t = callable.ReturnType()
	}

	return t, nil
}

func (self *CallExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitCallExpr(self)
}
