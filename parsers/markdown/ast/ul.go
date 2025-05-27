package ast

import (
	"fmt"
	"reflect"
)

type Ul struct {
	Content []Li
}

func (self *Ul) Add(items ...Li) *Ul {
	self.Content = append(self.Content, items...)
	return self
}

func (self Ul) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<ul>%s</ul>", string(content))
	return reflect.ValueOf(value), nil
}
