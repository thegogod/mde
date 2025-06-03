package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseLink(iter *tokens.Iterator) (core.Node, error) {
	link := html.A()

	for !iter.Match(tokens.RightBracket) {
		node, err := self.ParseInline(iter)

		if node == nil || err != nil {
			return link, err
		}

		link.Push(node)
	}

	if _, err := iter.Consume(tokens.LeftParen, "expected '('"); err != nil {
		return link, err
	}

	node, err := self.parseTextUntil(iter, tokens.RightParen)

	if node == nil || err != nil {
		return link, err
	}

	link.Href(string(node))
	return link, nil
}
