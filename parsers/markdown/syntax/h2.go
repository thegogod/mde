package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type H2 struct{}

func (self H2) IsBlock() bool {
	return true
}

func (self H2) IsInline() bool {
	return false
}

func (self H2) Name() string {
	return "h2"
}

func (self H2) Select(parser core.Parser, iter *core.Iterator) bool {
	if !iter.MatchCount(tokens.Hash, 2) {
		return false
	}

	return iter.Match(tokens.Space)
}

func (self H2) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	heading := html.H2()

	for iter.Curr().Kind() != tokens.Eof && iter.Curr().Kind() != tokens.NewLine {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
