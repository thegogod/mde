package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/tokens"
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

func (self StrikeAlt) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.MatchCount(tokens.Tilde, 2)
}

func (self StrikeAlt) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.MatchCount(tokens.Tilde, 2) {
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
