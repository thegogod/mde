package ast

import (
	"github.com/thegogod/mde/reflect"
)

type VarExpr struct {
	Name *Token
	Type reflect.Type
}

func NewVarExpr(name *Token, _type reflect.Type) *VarExpr {
	return &VarExpr{
		Name: name,
		Type: _type,
	}
}

func (self *VarExpr) GetType() (reflect.Type, error) {
	return self.Type, nil
}

func (self *VarExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitVarExpr(self)
}
