package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type BlockQuote struct {
	Content []core.Node
}

func (self *BlockQuote) Add(items ...core.Node) *BlockQuote {
	self.Content = append(self.Content, items...)
	return self
}

func (self BlockQuote) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<blockquote>%s</blockquote>", content)
	return value, nil
}
