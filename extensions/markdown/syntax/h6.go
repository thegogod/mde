package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

func (self H6) Select(parser core.Parser, iter *core.Iterator) bool {
	if !iter.MatchCount(core.Hash, 6) {
		return false
	}

	return iter.Match(core.Space)
}

func (self H6) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	heading := html.H6()

	for iter.Curr().Kind() != core.Eof && iter.Curr().Kind() != core.NewLine {
		node, err := parser.ParseInline(parser, iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
