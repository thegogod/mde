package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Italic struct {
	Content []core.Node
}

func (self *Italic) Add(items ...core.Node) *Italic {
	self.Content = append(self.Content, items...)
	return self
}

func (self Italic) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<i>%s</i>", content)
	return value, nil
}
