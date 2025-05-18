package ast

import (
	"github.com/thegogod/mde/reflect"
)

type GroupingExpr struct {
	Expr Expr
}

func NewGroupingExpr(expr Expr) *GroupingExpr {
	return &GroupingExpr{
		Expr: expr,
	}
}

func (self *GroupingExpr) GetType() (reflect.Type, error) {
	return self.Expr.GetType()
}

func (self *GroupingExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitGroupingExpr(self)
}
