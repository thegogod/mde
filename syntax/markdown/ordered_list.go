package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/tokens"
)

type OrderedList struct{}

func (self OrderedList) IsBlock() bool {
	return true
}

func (self OrderedList) IsInline() bool {
	return false
}

func (self OrderedList) Name() string {
	return "ordered_list"
}

func (self OrderedList) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.Match(tokens.Integer) && iter.Match(tokens.Period) && iter.Match(tokens.Space)
}

func (self OrderedList) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	it := iter.(*tokens.Iterator)
	it.ListDepth++
	ol := html.Ol()

	for {
		node, err := parser.ParseSyntax("list_item", iter)

		if node == nil || err != nil {
			it.ListDepth--
			return ol, err
		}

		ol.Push(node.(*html.ListItemElement))

		if !iter.MatchCount(tokens.Tab, it.ListDepth-1) {
			break
		}

		if !(iter.Match(tokens.Integer) && iter.Match(tokens.Period) && iter.Match(tokens.Space)) {
			break
		}
	}

	it.ListDepth--
	return ol, nil
}
