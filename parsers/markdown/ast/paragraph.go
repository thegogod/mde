package ast

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type Paragraph struct {
	Content []core.Node
}

func (self *Paragraph) Add(items ...core.Node) *Paragraph {
	self.Content = append(self.Content, items...)
	return self
}

func (self Paragraph) Render() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<p>%s</p>", bytes.TrimSpace(content))
	return reflect.ValueOf(value), nil
}
