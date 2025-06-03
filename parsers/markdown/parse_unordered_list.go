package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseUnorderedList(iter *tokens.Iterator) (core.Node, error) {
	iter.ListDepth++
	ul := html.Ul()

	for {
		node, err := self.parseListItem(iter)

		if node == nil || err != nil {
			iter.ListDepth--
			return ul, err
		}

		ul.Push(node)

		if !iter.MatchCount(tokens.Tab, iter.ListDepth-1) {
			break
		}

		if !iter.Match(tokens.Ul) {
			break
		}
	}

	iter.ListDepth--
	return ul, nil
}
