package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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
	return iter.Match(core.Bang)
}

func (self Image) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	image := html.Img()

	if _, err := iter.Consume(core.LeftBracket, "expected '['"); err != nil {
		return image, err
	}

	node, err := parser.ParseTextUntil(core.RightBracket, parser, iter)

	if node == nil || err != nil {
		return image, err
	}

	image.Alt(string(node))

	if _, err := iter.Consume(core.LeftParen, "expected '('"); err != nil {
		return image, err
	}

	node, err = parser.ParseTextUntil(core.RightParen, parser, iter)

	if node == nil || err != nil {
		return image, err
	}

	image.Src(string(node))
	return image, nil
}
