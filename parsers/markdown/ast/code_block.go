package ast

import (
	"fmt"

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

func (self CodeBlock) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	code := fmt.Appendf(nil, "<code>%s</code>", content)

	if self.Lang != nil {
		lang, err := self.Lang.Render()

		if err != nil {
			return []byte{}, err
		}

		code = fmt.Appendf(nil, `<code class="language-%s">%s</code>`, lang, content)
	}

	value := fmt.Appendf(nil, "<pre>%s</pre>", code)
	return value, nil
}
