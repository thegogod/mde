package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/tokens"
)

type NewLine struct{}

func (self NewLine) IsBlock() bool {
	return false
}

func (self NewLine) IsInline() bool {
	return true
}

func (self NewLine) Name() string {
	return "newline"
}

func (self NewLine) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.Match(tokens.NewLine)
}

func (self NewLine) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	if iter.Match(tokens.NewLine) {
		return nil, nil
	}

	it := iter.(*tokens.Iterator)

	for range it.BlockQuoteDepth {
		if !iter.Match(tokens.GreaterThan) {
			break
		}
	}

	curr := iter.Curr().String()

	if curr == " " || curr == "\n" {
		return nil, nil
	}

	return html.Raw("\n"), nil
}
