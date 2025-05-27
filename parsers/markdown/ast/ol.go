package ast

import (
	"fmt"
)

type Ol struct {
	Content []Li
}

func (self *Ol) Add(items ...Li) *Ol {
	self.Content = append(self.Content, items...)
	return self
}

func (self Ol) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<ol>%s</ol>", content)
	return value, nil
}
