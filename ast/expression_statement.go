package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type ExpressionStatement struct {
	expression Expression
}

func Expr(expression Expression) ExpressionStatement {
	return ExpressionStatement{expression}
}

func (self ExpressionStatement) Validate() error {
	return self.expression.Validate()
}

func (self ExpressionStatement) Evaluate(scope core.Scope) (reflect.Value, error) {
	_, err := self.expression.Evaluate(scope)
	return reflect.NewNil(), err
}
