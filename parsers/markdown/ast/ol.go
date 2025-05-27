package ast

import (
	"fmt"
	"reflect"
)

type Ol struct {
	Content []Li
}

func (self *Ol) Add(items ...Li) *Ol {
	self.Content = append(self.Content, items...)
	return self
}

func (self Ol) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<ol>%s</ol>", string(content))
	return reflect.ValueOf(value), nil
}
