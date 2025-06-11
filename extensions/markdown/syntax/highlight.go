package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Highlight struct{}

func (self Highlight) IsBlock() bool {
	return false
}

func (self Highlight) IsInline() bool {
	return true
}

func (self Highlight) Name() string {
	return "highlight"
}

func (self Highlight) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Equals, 2)
}

func (self Highlight) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	mark := html.Mark()

	for !iter.MatchCount(core.Equals, 2) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			return mark, iter.Curr().Error("expected closing '=='")
		}

		if err != nil {
			return mark, err
		}

		mark.Push(node)
	}

	return mark, nil
}
