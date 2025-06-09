package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/tokens"
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

func (self Break) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.MatchCount(tokens.Space, 2) && iter.Curr().Kind() == tokens.NewLine
}

func (self Break) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	return html.Br(), nil
}
