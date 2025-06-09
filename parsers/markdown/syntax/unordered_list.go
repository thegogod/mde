package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type UnOrderedList struct{}

func (self UnOrderedList) IsBlock() bool {
	return true
}

func (self UnOrderedList) IsInline() bool {
	return false
}

func (self UnOrderedList) Name() string {
	return "unordered_list"
}

func (self UnOrderedList) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(tokens.Dash) && iter.Match(tokens.Space)
}

func (self UnOrderedList) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	iter.ListDepth++
	ul := html.Ul()

	for {
		node, err := parser.ParseSyntax("list_item", parser, iter)

		if node == nil || err != nil {
			iter.ListDepth--
			return ul, err
		}

		ul.Push(node.(*html.ListItemElement))

		if !iter.MatchCount(tokens.Tab, iter.ListDepth-1) {
			break
		}

		if !(iter.Match(tokens.Dash) && iter.Match(tokens.Space)) {
			break
		}
	}

	iter.ListDepth--
	return ul, nil
}
