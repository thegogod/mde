package ast

import (
	"github.com/thegogod/mde/reflect"
)

type LogicalExpr struct {
	Left  Expr
	Op    *Token
	Right Expr
}

func NewLogicalExpr(left Expr, op *Token, right Expr) *LogicalExpr {
	return &LogicalExpr{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (self *LogicalExpr) GetType() (reflect.Type, error) {
	left, err := self.Left.GetType()

	if err != nil {
		return nil, err
	}

	right, err := self.Right.GetType()

	if err != nil {
		return nil, err
	}

	if !left.Equals(right) {
		return nil, NewError(
			self.Op.Path,
			self.Op.Ln,
			self.Op.Start,
			self.Op.End,
			"expected type '"+left.Name()+"', received '"+right.Name()+"'",
		)
	}

	return left, nil
}

func (self *LogicalExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitLogicalExpr(self)
}
