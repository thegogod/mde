package ast

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/reflect"
)

type BlockStatement struct {
	statements []Statement
}

func Block(statements ...Statement) BlockStatement {
	return BlockStatement{statements}
}

func (self BlockStatement) Validate() error {
	for _, statement := range self.statements {
		if err := statement.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (self BlockStatement) Evaluate(scope core.Scope) (reflect.Value, error) {
	child := scope.Create()

	for _, statement := range self.statements {
		value, err := statement.Evaluate(child)

		if err != nil {
			return value, err
		}

		if !value.IsNil() {
			return value, nil
		}
	}

	return reflect.NewNil(), nil
}
