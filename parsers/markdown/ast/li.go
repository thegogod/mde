package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Li struct {
	Content []core.Node
}

func (self *Li) Add(items ...core.Node) *Li {
	self.Content = append(self.Content, items...)
	return self
}

func (self Li) Render() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, bytes.TrimSpace(value.Bytes())...)
	}

	value := fmt.Appendf(nil, "<li>%s</li>", content)
	return reflect.ValueOf(value), nil
}
