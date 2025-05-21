package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Bold struct {
	Content []core.Node
}

func (self *Bold) Add(items ...core.Node) *Bold {
	self.Content = append(self.Content, items...)
	return self
}

func (self Bold) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<b>%s</b>", string(content))
	return reflect.ValueOf(value), nil
}
