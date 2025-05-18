package ast

import (
	"github.com/thegogod/mde/reflect"
)

type SelfExpr struct {
	Keyword *Token
	Type    reflect.Type
}

func NewSelfExpr(keyword *Token, _type reflect.Type) *SelfExpr {
	return &SelfExpr{
		Keyword: keyword,
		Type:    _type,
	}
}

func (self *SelfExpr) GetType() (reflect.Type, error) {
	return self.Type, nil
}

func (self *SelfExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitSelfExpr(self)
}
