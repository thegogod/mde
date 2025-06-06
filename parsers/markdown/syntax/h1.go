package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type H1 struct{}

func (self H1) IsBlock() bool {
	return true
}

func (self H1) IsInline() bool {
	return false
}

func (self H1) Name() string {
	return "h1"
}

func (self H1) Select(parser core.Parser, iter core.Iterator) bool {
	if !iter.MatchCount(tokens.Hash, 1) {
		return false
	}

	return iter.Match(tokens.Space)
}

func (self H1) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	heading := html.H1()

	for iter.Curr().Kind() != tokens.Eof && iter.Curr().Kind() != tokens.NewLine {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
