package ast

import (
	"fmt"
)

type Ul struct {
	Content []Li
}

func (self *Ul) Add(items ...Li) *Ul {
	self.Content = append(self.Content, items...)
	return self
}

func (self Ul) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<ul>%s</ul>", content)
	return value, nil
}
