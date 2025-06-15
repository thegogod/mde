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
		return self.op.Error(fmt.Sprintf("expected type '%s', received '%s'", left.Name(), right.Name()))
	}

	return nil
}

func (self BinaryExpression) Evaluate(scope core.Scope) (reflect.Value, error) {
	left, err := self.left.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	right, err := self.right.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	switch self.op.Kind() {
	case core.EqualsEquals:
		return reflect.NewBool(left.Eq(right)), nil
	case core.NotEquals:
		return reflect.NewBool(!left.Eq(right)), nil
	case core.GreaterThan:
		return reflect.NewBool(left.Gt(right)), nil
	case core.GreaterThanEquals:
		return reflect.NewBool(left.GtEq(right)), nil
	case core.LessThan:
		return reflect.NewBool(left.Lt(right)), nil
	case core.LessThanEquals:
		return reflect.NewBool(left.LtEq(right)), nil
	case core.Plus:
		if left.Numeric() {
			return left.Add(right), nil
		}

		return reflect.NewString(left.String() + right.String()), nil
	case core.Dash:
		return left.Subtract(right), nil
	case core.Asterisk:
		return left.Multiply(right), nil
	case core.Slash:
		return left.Divide(right), nil
	}

	return reflect.NewNil(), nil
}
