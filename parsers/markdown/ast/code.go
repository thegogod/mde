package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Code struct {
	Content []core.Node
}

func (self *Code) Add(items ...core.Node) *Code {
	self.Content = append(self.Content, items...)
	return self
}

func (self Code) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<code>%s</code>", string(content))
	return reflect.ValueOf(value), nil
}
