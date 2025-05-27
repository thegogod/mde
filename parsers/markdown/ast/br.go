package ast

import (
	"reflect"
)

type Br struct{}

func (self Br) Render() (reflect.Value, error) {
	return reflect.ValueOf([]byte("<br>")), nil
}
