package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type H3 struct{}

func (self H3) IsBlock() bool {
	return true
}

func (self H3) IsInline() bool {
	return false
}

func (self H3) Name() string {
	return "h3"
}

func (self H3) Select(parser core.Parser, iter core.Iterator) bool {
	if !iter.MatchCount(tokens.Hash, 3) {
		return false
	}

	return iter.Match(tokens.Space)
}

func (self H3) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	heading := html.H3()

	for iter.Curr().IsInline() {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
