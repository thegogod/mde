package ast

import (
	"reflect"

	"github.com/thegogod/mde/core"
)

type Group struct {
	Items []core.Node
}

func (self *Group) Add(items ...core.Node) *Group {
	self.Items = append(self.Items, items...)
	return self
}

func (self Group) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Items {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	return reflect.ValueOf(content), nil
}
