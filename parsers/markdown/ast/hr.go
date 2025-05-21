package ast

import (
	"reflect"
)

type Hr struct{}

func (self Hr) Eval() (reflect.Value, error) {
	return reflect.ValueOf([]byte("<hr>")), nil
}
