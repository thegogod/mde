package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Italic struct{}

func (self Italic) IsBlock() bool {
	return false
}

func (self Italic) IsInline() bool {
	return true
}

func (self Italic) Name() string {
	return "italic"
}

func (self Italic) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(tokens.Asterisk, 1)
}

func (self Italic) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	italic := html.I()

	for !iter.Match(tokens.Asterisk) {
		node, err := parser.ParseInline(parser, iter)

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
