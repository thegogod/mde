package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type CodeBlock struct {
	Content []core.Node
}

func (self *CodeBlock) Add(items ...core.Node) *CodeBlock {
	self.Content = append(self.Content, items...)
	return self
}

func (self CodeBlock) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<pre><code>%s</code></pre>", string(content))
	return reflect.ValueOf(value), nil
}
