package ast

import (
	"github.com/thegogod/mde/reflect"
)

type GetExpr struct {
	Object Expr
	Name   *Token
}

func NewGetExpr(object Expr, name *Token) *GetExpr {
	return &GetExpr{
		Object: object,
		Name:   name,
	}
}

func (self *GetExpr) GetType() (reflect.Type, error) {
	t, err := self.Object.GetType()

	if err != nil {
		return nil, err
	}

	return t.GetMember(self.Name.String()), nil
}

func (self *GetExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitGetExpr(self)
}
