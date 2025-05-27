package ast

import (
	"reflect"

	"github.com/thegogod/mde/core"
)

type Text struct {
	Content core.Token
}

func (self Text) Render() (reflect.Value, error) {
	return reflect.ValueOf(self.Content.GetBytes()), nil
}
