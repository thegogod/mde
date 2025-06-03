package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type StrikeAlt struct{}

func (self StrikeAlt) IsBlock() bool {
	return false
}

func (self StrikeAlt) IsInline() bool {
	return true
}

func (self StrikeAlt) Name() string {
	return "strike_alt"
}

func (self StrikeAlt) Select(iter core.Iterator) bool {
	return iter.Match(tokens.StrikeAlt)
}

func (self StrikeAlt) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(tokens.StrikeAlt) {
		node, err := parser.ParseInline(iter)

		if node == nil {
			return strike, iter.Curr().Error("expected closing '~~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
