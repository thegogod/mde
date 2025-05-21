package core

import "reflect"

type Node interface {
	Eval() (reflect.Value, error)
}
