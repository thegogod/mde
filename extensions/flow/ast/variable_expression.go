package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type VariableExpression struct {
	name  core.Token
	_type reflect.Type
}

func VarRef(name core.Token, _type reflect.Type) VariableExpression {
	return VariableExpression{name, _type}
}

func (self VariableExpression) Type() reflect.Type {
	return self._type
}

func (self VariableExpression) Validate() error {
	return nil
}

func (self VariableExpression) Evaluate(scope core.Scope) (reflect.Value, error) {
	if !scope.Has(self.name.String()) {
		return reflect.NewNil(), fmt.Errorf("identifier '%s' not found", self.name.String())
	}

	return scope.Get(self.name.String())
}
