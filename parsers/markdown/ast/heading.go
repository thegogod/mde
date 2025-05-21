package ast

import (
	"fmt"
	"reflect"

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

func (self Heading) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<h%d>%s</h%d>", self.Depth, string(content), self.Depth)
	return reflect.ValueOf(value), nil
}
