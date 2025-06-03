package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
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

func (self Link) Select(iter core.Iterator) bool {
	return iter.Match(tokens.Link)
}

func (self Link) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	link := html.A()

	for !iter.Match(tokens.RightBracket) {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return link, err
		}

		link.Push(node)
	}

	if _, err := iter.Consume(tokens.LeftParen, "expected '('"); err != nil {
		return link, err
	}

	node, err := parser.ParseTextUntil(iter, tokens.RightParen)

	if node == nil || err != nil {
		return link, err
	}

	link.Href(string(node))
	return link, nil
}
