package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Bold struct {
	Content []core.Node
}

func (self *Bold) Add(items ...core.Node) *Bold {
	self.Content = append(self.Content, items...)
	return self
}

func (self Bold) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<strong>%s</strong>", content)
	return value, nil
}
