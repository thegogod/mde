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

func (self Bold) Render() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<strong>%s</strong>", content)
	return reflect.ValueOf(value), nil
}
