package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type ItalicAlt struct{}

func (self ItalicAlt) IsBlock() bool {
	return false
}

func (self ItalicAlt) IsInline() bool {
	return true
}

func (self ItalicAlt) Name() string {
	return "italic_alt"
}

func (self ItalicAlt) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Underscore, 1)
}

func (self ItalicAlt) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	italic := html.I()

	for !iter.Match(core.Underscore) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			return italic, iter.Curr().Error("expected closing '_'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}
