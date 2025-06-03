package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseImage(iter *tokens.Iterator) (core.Node, error) {
	image := html.Img()

	if _, err := iter.Consume(tokens.LeftBracket, "expected '['"); err != nil {
		return image, err
	}

	node, err := self.parseTextUntil(iter, tokens.RightBracket)

	if node == nil || err != nil {
		return image, err
	}

	image.Alt(string(node))

	if _, err := iter.Consume(tokens.LeftParen, "expected '('"); err != nil {
		return image, err
	}

	node, err = self.parseTextUntil(iter, tokens.RightParen)

	if node == nil || err != nil {
		return image, err
	}

	image.Src(string(node))
	return image, nil
}
