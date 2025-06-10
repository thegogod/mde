package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Link struct{}

func (self Link) IsBlock() bool {
	return false
}

func (self Link) IsInline() bool {
	return true
}

func (self Link) Name() string {
	return "link"
}

func (self Link) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(core.LeftBracket)
}

func (self Link) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	link := html.A()

	for !iter.Match(core.RightBracket) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil || err != nil {
			return link, err
		}

		link.Push(node)
	}

	if _, err := iter.Consume(core.LeftParen, "expected '('"); err != nil {
		return link, err
	}

	node, err := parser.ParseTextUntil(core.RightParen, parser, iter)

	if node == nil || err != nil {
		return link, err
	}

	link.Href(string(node))
	return link, nil
}
