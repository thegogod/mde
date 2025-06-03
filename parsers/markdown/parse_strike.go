package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseStrike(iter *tokens.Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(tokens.Strike) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return strike, iter.Curr().Error("expected closing '~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}

func (self *Parser) parseStrikeAlt(iter *tokens.Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(tokens.StrikeAlt) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return strike, iter.Curr().Error("expected closing '~~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
