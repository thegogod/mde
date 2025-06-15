package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type LiteralExpression struct {
	value reflect.Value
}

func Literal(value reflect.Value) LiteralExpression {
	return LiteralExpression{value: value}
}

func (self LiteralExpression) Type() reflect.Type {
	return self.value.Type()
}

func (self LiteralExpression) Validate() error {
	return nil
}

func (self LiteralExpression) Evaluate(scope core.Scope) (reflect.Value, error) {
	return self.value, nil
}
