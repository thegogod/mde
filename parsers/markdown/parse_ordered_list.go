package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseOrderedList(iter *tokens.Iterator) (core.Node, error) {
	iter.ListDepth++
	ol := html.Ol()

	for {
		node, err := self.parseListItem(iter)

		if node == nil || err != nil {
			iter.ListDepth--
			return ol, err
		}

		ol.Push(node)

		if !iter.MatchCount(tokens.Tab, iter.ListDepth-1) {
			break
		}

		if !iter.Match(tokens.Ol) {
			break
		}
	}

	iter.ListDepth--
	return ol, nil
}
