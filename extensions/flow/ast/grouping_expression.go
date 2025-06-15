package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type GroupingExpression struct {
	expr Expression
}

func Grouping(expr Expression) GroupingExpression {
	return GroupingExpression{expr}
}

func (self GroupingExpression) Type() reflect.Type {
	return self.expr.Type()
}

func (self GroupingExpression) Validate() error {
	return self.expr.Validate()
}

func (self GroupingExpression) Evaluate(scope core.Scope) (reflect.Value, error) {
	return self.expr.Evaluate(scope)
}
