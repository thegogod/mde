package core

import "reflect"

type Statement interface {
	Type() (reflect.Type, error)
	Evaluate()
}

type Expression interface {
	Type() (reflect.Type, error)
	Evaluate()
}

type Interpreter interface {
	OnStatement(stmt Statement) (reflect.Value, error)
	OnExpression(expr Expression) (reflect.Value, error)
}
