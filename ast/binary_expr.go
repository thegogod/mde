package ast

import (
	"github.com/thegogod/mde/reflect"
)

type BinaryExpr struct {
	Left  Expr
	Op    *Token
	Right Expr
}

func NewBinaryExpr(left Expr, op *Token, right Expr) *BinaryExpr {
	return &BinaryExpr{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (self *BinaryExpr) GetType() (reflect.Type, error) {
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

func (self *BinaryExpr) Accept(v ExprVisitor) (*reflect.Value, error) {
	return v.VisitBinaryExpr(self)
}
