package ast

import (
	"fmt"
	"reflect"

	"github.com/thegogod/mde/core"
)

type BlockQuote struct {
	Content []core.Node
}

func (self *BlockQuote) Add(items ...core.Node) *BlockQuote {
	self.Content = append(self.Content, items...)
	return self
}

func (self BlockQuote) Eval() (reflect.Value, error) {
	content := []byte{}

	for _, item := range self.Content {
		value, err := item.Eval()

		if err != nil {
			return reflect.Value{}, err
		}

		content = append(content, value.Bytes()...)
	}

	value := fmt.Appendf(nil, "<blockquote>%s</blockquote>", string(content))
	return reflect.ValueOf(value), nil
}
