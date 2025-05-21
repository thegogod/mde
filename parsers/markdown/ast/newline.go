package ast

import (
	"reflect"
)

type NewLine struct{}

func (self NewLine) Eval() (reflect.Value, error) {
	return reflect.ValueOf([]byte(" ")), nil
}
