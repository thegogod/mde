package ast

import (
	"fmt"

	"github.com/thegogod/mde/core"
)

type Heading struct {
	Depth   int
	Content []core.Node
}

func (self *Heading) Add(items ...core.Node) *Heading {
	self.Content = append(self.Content, items...)
	return self
}

func (self Heading) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	value := fmt.Appendf(nil, "<h%d>%s</h%d>", self.Depth, content, self.Depth)
	return value, nil
}
