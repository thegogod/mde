package ast

import (
	"reflect"
)

type Br struct{}

func (self Br) Eval() (reflect.Value, error) {
	return reflect.ValueOf([]byte("<br>")), nil
}
