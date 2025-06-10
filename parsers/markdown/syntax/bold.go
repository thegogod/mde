package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

type Bold struct{}

func (self Bold) IsBlock() bool {
	return false
}

func (self Bold) IsInline() bool {
	return true
}

func (self Bold) Name() string {
	return "bold"
}

func (self Bold) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Asterisk, 2)
}

func (self Bold) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	bold := html.Strong()

	for !iter.MatchCount(core.Asterisk, 2) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			return bold, iter.Curr().Error("expected closing '**'")
		}

		if err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}
