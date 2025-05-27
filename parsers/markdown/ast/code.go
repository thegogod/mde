package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Code struct {
	Content []core.Node
}

func (self *Code) Add(items ...core.Node) *Code {
	self.Content = append(self.Content, items...)
	return self
}

func (self Code) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<code>%s</code>", content)
	return value, nil
}
