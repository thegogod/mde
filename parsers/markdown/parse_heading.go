package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseHeading(depth int, iter *tokens.Iterator) (core.Node, error) {
	heading := html.Heading(depth)

	for iter.Curr().IsInline() {
		node, err := self.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
