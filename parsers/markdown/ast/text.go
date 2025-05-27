package ast

import (
	"github.com/thegogod/mde/core"
)

type Text struct {
	Content core.Token
}

func (self Text) Render() ([]byte, error) {
	return self.Content.GetBytes(), nil
}
