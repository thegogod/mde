package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type BlockQuote struct{}

func (self BlockQuote) IsBlock() bool {
	return true
}

func (self BlockQuote) IsInline() bool {
	return false
}

func (self BlockQuote) Name() string {
	return "block_quote"
}

func (self BlockQuote) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(tokens.GreaterThan)
}

func (self BlockQuote) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	iter.BlockQuoteDepth++
	blockQuote := html.BlockQuote()

	for {
		node, err := parser.ParseBlock(iter)

		if node == nil || err != nil {
			iter.BlockQuoteDepth--
			return blockQuote, err
		}

		blockQuote.Push(node)

		if iter.Curr().Kind() != tokens.GreaterThan {
			break
		}
	}

	iter.BlockQuoteDepth--
	return blockQuote, nil
}
