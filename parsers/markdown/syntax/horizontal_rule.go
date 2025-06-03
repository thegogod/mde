package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type HorizontalRule struct{}

func (self HorizontalRule) IsBlock() bool {
	return true
}

func (self HorizontalRule) IsInline() bool {
	return false
}

func (self HorizontalRule) Name() string {
	return "hr"
}

func (self HorizontalRule) Select(iter core.Iterator) bool {
	return iter.Match(tokens.Hr)
}

func (self HorizontalRule) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	return html.Hr(), nil
}
