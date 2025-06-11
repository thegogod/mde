package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

func (self OrderedList) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(core.Integer) && iter.Match(core.Period) && iter.Match(core.Space)
}

func (self OrderedList) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	iter.ListDepth++
	ol := html.Ol()

	for {
		node, err := parser.ParseSyntax("markdown", "list_item", parser, iter)

		if node == nil || err != nil {
			iter.ListDepth--
			return ol, err
		}

		ol.Push(node.(*html.ListItemElement))

		if !iter.MatchCount(core.Tab, iter.ListDepth-1) {
			break
		}

		if !(iter.Match(core.Integer) && iter.Match(core.Period) && iter.Match(core.Space)) {
			break
		}
	}

	iter.ListDepth--
	return ol, nil
}
