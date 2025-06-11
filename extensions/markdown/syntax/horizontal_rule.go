package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

func (self HorizontalRule) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.Dash, 3)
}

func (self HorizontalRule) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	return html.Hr(), nil
}
