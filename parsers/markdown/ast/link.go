package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Link struct {
	Text []core.Node
	Href []core.Node
}

func (self Link) Eval() (reflect.Value, error) {
	text := []byte{}

	for _, item := range self.Text {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		text = append(text, value.Bytes()...)
	}

	href := []byte{}

	for _, item := range self.Href {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		href = append(href, value.Bytes()...)
	}

	value := fmt.Appendf(nil, `<a href="%s">%s</a>`, href, text)
	return reflect.ValueOf(value), nil
}
