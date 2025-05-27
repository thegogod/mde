package ast

import (
	"reflect"
)

type NewLine struct{}

func (self NewLine) Render() (reflect.Value, error) {
	return reflect.ValueOf([]byte("\n")), nil
}
