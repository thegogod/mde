package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type UnaryExpression struct {
	op    core.Token
	right Expression
}

func Unary(op core.Token, right Expression) UnaryExpression {
	return UnaryExpression{
		op:    op,
		right: right,
	}
}

func (self UnaryExpression) Type() reflect.Type {
	return self.right.Type()
}

func (self UnaryExpression) Validate() error {
	return self.right.Validate()
}

func (self UnaryExpression) Evaluate(scope core.Scope) (reflect.Value, error) {
	right, err := self.right.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	switch self.op.Kind() {
	case core.Bang:
		return reflect.NewBool(!right.Truthy()), nil
	case core.Dash:
		right.Decrement()
		return right, nil
	}

	return right, nil
}
