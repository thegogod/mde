package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Strike struct{}

func (self Strike) IsBlock() bool {
	return false
}

func (self Strike) IsInline() bool {
	return true
}

func (self Strike) Name() string {
	return "strike"
}

func (self Strike) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Tilde, 1)
}

func (self Strike) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(core.Tilde) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			return strike, iter.Curr().Error("expected closing '~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
