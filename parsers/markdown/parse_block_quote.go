package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseBlockQuote(iter *tokens.Iterator) (core.Node, error) {
	iter.BlockQuoteDepth++
	blockQuote := html.BlockQuote()

	for {
		node, err := self.ParseBlock(iter)

		if node == nil || err != nil {
			iter.BlockQuoteDepth--
			return blockQuote, err
		}

		blockQuote.Push(node)

		if iter.Curr().Kind() != tokens.BlockQuote {
			break
		}
	}

	iter.BlockQuoteDepth--
	return blockQuote, nil
}
