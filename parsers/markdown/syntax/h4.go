package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type H4 struct{}

func (self H4) IsBlock() bool {
	return true
}

func (self H4) IsInline() bool {
	return false
}

func (self H4) Name() string {
	return "h4"
}

func (self H4) Select(parser core.Parser, iter *core.Iterator) bool {
	if !iter.MatchCount(tokens.Hash, 4) {
		return false
	}

	return iter.Match(tokens.Space)
}

func (self H4) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	heading := html.H4()

	for iter.Curr().Kind() != tokens.Eof && iter.Curr().Kind() != tokens.NewLine {
		node, err := parser.ParseInline(parser, iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
