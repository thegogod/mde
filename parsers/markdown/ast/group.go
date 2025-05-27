package ast

import (
	"github.com/thegogod/mde/core"
)

type Group struct {
	Items []core.Node
}

func (self *Group) Add(items ...core.Node) *Group {
	self.Items = append(self.Items, items...)
	return self
}

func (self Group) Render() ([]byte, error) {
	content := []byte{}

	for _, item := range self.Items {
		value, err := item.Render()

		if err != nil {
			return []byte{}, err
		}

		content = append(content, value...)
	}

	return content, nil
}
