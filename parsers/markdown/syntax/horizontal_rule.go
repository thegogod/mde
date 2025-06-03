package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type HorizontalRule struct{}

func (self *HorizontalRule) IsBlock() bool {
	return true
}

func (self *HorizontalRule) IsInline() bool {
	return false
}

func (self *HorizontalRule) Select(iter core.Iterator) bool {
	return iter.Curr().Kind() == tokens.Hr
}

func (self *HorizontalRule) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	hr := html.Hr()

	if _, err := iter.Consume(tokens.Hr, "expected '---'"); err != nil {
		return hr, err
	}

	return hr, nil
}
