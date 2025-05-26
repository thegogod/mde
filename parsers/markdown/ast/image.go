package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Image struct {
	Alt []core.Node
	Src []core.Node
}

func (self Image) Eval() (reflect.Value, error) {
	alt := []byte{}

	for _, item := range self.Alt {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		alt = append(alt, value.Bytes()...)
	}

	src := []byte{}

	for _, item := range self.Src {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		src = append(src, value.Bytes()...)
	}

	value := fmt.Appendf(nil, `<img alt="%s" src="%s" />`, alt, src)
	return reflect.ValueOf(value), nil
}
