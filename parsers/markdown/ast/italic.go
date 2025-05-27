package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Italic struct {
	Content []core.Node
}

func (self *Italic) Add(items ...core.Node) *Italic {
	self.Content = append(self.Content, items...)
	return self
}

func (self Italic) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<i>%s</i>", content)
	return reflect.ValueOf(value), nil
}
