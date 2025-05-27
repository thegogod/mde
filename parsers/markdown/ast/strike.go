package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Strike struct {
	Content []core.Node
}

func (self *Strike) Add(items ...core.Node) *Strike {
	self.Content = append(self.Content, items...)
	return self
}

func (self Strike) Render() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<s>%s</s>", content)
	return reflect.ValueOf(value), nil
}
