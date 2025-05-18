package ast

import (
	"github.com/thegogod/mde/reflect"
)

type SetExpr struct {
	Object Expr
	Name   *Token
	Value  Expr
}

func NewSetExpr(object Expr, name *Token, value Expr) *SetExpr {
	return &SetExpr{
		Object: object,
		Name:   name,
		Value:  value,
	}
}

func (self *SetExpr) GetType() (reflect.Type, error) {
	object, err := self.Object.GetType()

	if err != nil {
		return nil, err
	}

	value, err := self.Value.GetType()

	if err != nil {
		return nil, err
	}

	if !object.Equals(value) {
		return nil, NewError(
			self.Name.Path,
			self.Name.Ln,
			self.Name.Start,
			self.Name.End,
			"expected type '"+object.Name()+"', received '"+value.Name()+"'",
		)
	}

	return object, nil
}

func (self *SetExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitSetExpr(self)
}
