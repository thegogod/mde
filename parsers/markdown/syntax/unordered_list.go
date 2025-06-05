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

func (self UnOrderedList) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.Match(tokens.Ul)
}

func (self UnOrderedList) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	it := iter.(*tokens.Iterator)
	it.ListDepth++
	ul := html.Ul()

	for {
		node, err := parser.ParseSyntax("list_item", iter)

		if node == nil || err != nil {
			it.ListDepth--
			return ul, err
		}

		ul.Push(node.(*html.ListItemElement))

		if !iter.MatchCount(tokens.Tab, it.ListDepth-1) {
			break
		}

		if !iter.Match(tokens.Ul) {
			break
		}
	}

	it.ListDepth--
	return ul, nil
}
