package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type BoldAlt struct{}

func (self BoldAlt) IsBlock() bool {
	return false
}

func (self BoldAlt) IsInline() bool {
	return true
}

func (self BoldAlt) Name() string {
	return "bold_alt"
}

func (self BoldAlt) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(tokens.Underscore, 2)
}

func (self BoldAlt) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	bold := html.Strong()

	for !iter.MatchCount(tokens.Underscore, 2) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			return bold, iter.Curr().Error("expected closing '__'")
		}

		if err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}
