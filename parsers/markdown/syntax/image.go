package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Image struct{}

func (self Image) IsBlock() bool {
	return false
}

func (self Image) IsInline() bool {
	return true
}

func (self Image) Name() string {
	return "image"
}

func (self Image) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(tokens.Bang)
}

func (self Image) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	image := html.Img()

	if _, err := iter.Consume(tokens.LeftBracket, "expected '['"); err != nil {
		return image, err
	}

	node, err := parser.ParseTextUntil(tokens.RightBracket, iter)

	if node == nil || err != nil {
		return image, err
	}

	image.Alt(string(node))

	if _, err := iter.Consume(tokens.LeftParen, "expected '('"); err != nil {
		return image, err
	}

	node, err = parser.ParseTextUntil(tokens.RightParen, iter)

	if node == nil || err != nil {
		return image, err
	}

	image.Src(string(node))
	return image, nil
}
