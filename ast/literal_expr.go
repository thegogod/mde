package ast

import (
	"github.com/thegogod/mde/reflect"
)

type LiteralExpr struct {
	Value *reflect.Value
}

func NewLiteralExpr(value *reflect.Value) *LiteralExpr {
	return &LiteralExpr{
		Value: value,
	}
}

func (self *LiteralExpr) GetType() (reflect.Type, error) {
	return self.Value.Type(), nil
}

func (self *LiteralExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitLiteralExpr(self)
}
