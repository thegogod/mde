package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

func (self H3) Select(parser core.Parser, iter *core.Iterator) bool {
	if !iter.MatchCount(core.Hash, 3) {
		return false
	}

	return iter.Match(core.Space)
}

func (self H3) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	heading := html.H3()

	for iter.Curr().Kind() != core.Eof && iter.Curr().Kind() != core.NewLine {
		node, err := parser.ParseInline(parser, iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
