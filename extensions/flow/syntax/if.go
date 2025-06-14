package syntax

import (
	"github.com/thegogod/mde/core"
)

type If struct{}

func (self If) IsBlock() bool {
	return true
}

func (self If) IsInline() bool {
	return true
}

func (self If) Name() string {
	return "if"
}

func (self If) Select(parser core.Parser, iter *core.Iterator) bool {
	if !iter.Match(core.At) || !iter.MatchBytes('i', 'f') {
		return false
	}

	iter.NextWhile(core.Space, core.Tab, core.NewLine)

	if !iter.Match(core.LeftBrace) {
		return false
	}

	return true
}

func (self If) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	return nil, nil
}
