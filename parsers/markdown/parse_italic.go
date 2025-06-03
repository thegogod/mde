package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseItalic(iter *tokens.Iterator) (core.Node, error) {
	italic := html.I()

	for !iter.Match(tokens.Italic) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return italic, iter.Curr().Error("expected closing '*'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}

func (self *Parser) parseItalicAlt(iter *tokens.Iterator) (core.Node, error) {
	italic := html.I()

	for !iter.Match(tokens.ItalicAlt) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return italic, iter.Curr().Error("expected closing '_'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}
