package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Heading struct{}

func (self Heading) IsBlock() bool {
	return true
}

func (self Heading) IsInline() bool {
	return false
}

func (self Heading) Name() string {
	return "heading"
}

func (self Heading) Select(iter core.Iterator) bool {
	return iter.Match(tokens.H1) ||
		iter.Match(tokens.H2) ||
		iter.Match(tokens.H3) ||
		iter.Match(tokens.H4) ||
		iter.Match(tokens.H5) ||
		iter.Match(tokens.H6)
}

func (self Heading) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	depth := len(iter.Prev().Bytes()) - 1
	heading := html.Heading(depth)

	for iter.Curr().IsInline() {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
