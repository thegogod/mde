package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type BinaryExpression struct {
	left  Expression
	op    core.Token
	right Expression
}

func Binary(left Expression, op core.Token, right Expression) BinaryExpression {
	return BinaryExpression{
		left:  left,
		op:    op,
		right: right,
	}
}

func (self BinaryExpression) Type() reflect.Type {
	return self.left.Type()
}

func (self BinaryExpression) Validate() error {
	if err := self.left.Validate(); err != nil {
		return err
	}

	if err := self.right.Validate(); err != nil {
		return err
	}

	left := self.left.Type()
	right := self.right.Type()

	if !left.Equals(right) {
		return core.Err(
			self.op.Start(),
			self.op.End(),
			fmt.Sprintf("expected type '%s', received '%s'", left.Name(), right.Name()),
		)
	}

	return nil
}

func (self BinaryExpression) Evaluate(scope core.Scope) (reflect.Value, error) {
	_, err := self.left.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	_, err = self.right.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	return reflect.NewNil(), nil
}
