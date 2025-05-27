package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type CodeBlock struct {
	Lang    core.Node
	Content []core.Node
}

func (self *CodeBlock) Add(items ...core.Node) *CodeBlock {
	self.Content = append(self.Content, items...)
	return self
}

func (self CodeBlock) Render() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	code := fmt.Appendf(nil, "<code>%s</code>", content)

	if self.Lang != nil {
		lang, err := self.Lang.Render()

		if err != nil {
			return reflect.Value{}, err
		}

		code = fmt.Appendf(nil, `<code class="language-%s">%s</code>`, lang.Bytes(), content)
	}

	value := fmt.Appendf(nil, "<pre>%s</pre>", code)
	return reflect.ValueOf(value), nil
}
