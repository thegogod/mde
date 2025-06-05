package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type H5 struct{}

func (self H5) IsBlock() bool {
	return true
}

func (self H5) IsInline() bool {
	return false
}

func (self H5) Name() string {
	return "h5"
}

func (self H5) Select(parser core.Parser, iter core.Iterator) bool {
	if !iter.MatchCount(tokens.Hash, 5) {
		return false
	}

	return iter.Match(tokens.Space)
}

func (self H5) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	heading := html.H5()

	for iter.Curr().IsInline() {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
