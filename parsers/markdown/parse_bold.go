package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseBold(iter *tokens.Iterator) (core.Node, error) {
	bold := html.Strong()

	for !iter.Match(tokens.Bold) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return bold, iter.Curr().Error("expected closing '**'")
		}

		if err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}

func (self *Parser) parseBoldAlt(iter *tokens.Iterator) (core.Node, error) {
	bold := html.Strong()

	for !iter.Match(tokens.BoldAlt) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return bold, iter.Curr().Error("expected closing '__'")
		}

		if err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}
