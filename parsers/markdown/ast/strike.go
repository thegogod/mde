package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Strike struct {
	Content []core.Node
}

func (self *Strike) Add(items ...core.Node) *Strike {
	self.Content = append(self.Content, items...)
	return self
}

func (self Strike) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<s>%s</s>", content)
	return value, nil
}
