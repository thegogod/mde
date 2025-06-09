package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/tokens"
)

type H6 struct{}

func (self H6) IsBlock() bool {
	return true
}

func (self H6) IsInline() bool {
	return false
}

func (self H6) Name() string {
	return "h6"
}

func (self H6) Select(parser core.Parser, iter core.Iterator) bool {
	if !iter.MatchCount(tokens.Hash, 6) {
		return false
	}

	return iter.Match(tokens.Space)
}

func (self H6) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	heading := html.H6()

	for iter.Curr().Kind() != tokens.Eof && iter.Curr().Kind() != tokens.NewLine {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
