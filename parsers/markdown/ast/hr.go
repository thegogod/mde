package ast

import (
	"reflect"
)

type Hr struct{}

func (self Hr) Render() (reflect.Value, error) {
	return reflect.ValueOf([]byte("<hr>")), nil
}
