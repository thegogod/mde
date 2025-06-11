package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Break struct{}

func (self Break) IsBlock() bool {
	return false
}

func (self Break) IsInline() bool {
	return true
}

func (self Break) Name() string {
	return "br"
}

func (self Break) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Space, 2) && iter.Curr().Kind() == core.NewLine
}

func (self Break) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	return html.Br(), nil
}
