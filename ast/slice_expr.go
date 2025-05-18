package ast

import (
	"github.com/thegogod/mde/reflect"
)

type SliceExpr struct {
	Type  reflect.Type
	Items []Expr
}

func NewSliceExpr(_type reflect.Type, items []Expr) *SliceExpr {
	return &SliceExpr{
		Type:  reflect.NewSliceType(_type, -1),
		Items: items,
	}
}

func (self *SliceExpr) GetType() (reflect.Type, error) {
	return self.Type, nil
}

func (self *SliceExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitSliceExpr(self)
}
