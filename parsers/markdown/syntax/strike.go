package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
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

func (self Strike) Select(iter core.Iterator) bool {
	return iter.Match(tokens.Strike)
}

func (self Strike) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(tokens.Strike) {
		node, err := parser.ParseInline(iter)

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
