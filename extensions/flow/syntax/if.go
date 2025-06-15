package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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
	if !iter.Match(core.At) || !iter.MatchLiteral("if") {
		return false
	}

	return true
}

func (self If) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	body := html.Fragment()
	iter.NextWhile(core.Space, core.Tab, core.NewLine)

	if !iter.Match(core.LeftParen) {
		return body, iter.Curr().Error("expected '('")
	}

	iter.NextWhile(core.Space, core.Tab, core.NewLine)

	if !iter.Match(core.LeftBrace) {
		return body, iter.Curr().Error("expected '{'")
	}

	iter.NextWhile(core.Space, core.Tab, core.NewLine)

	for !iter.Match(core.RightBrace) {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			return body, iter.Curr().Error("expected closing '}'")
		}

		if err != nil {
			return body, err
		}

		body.Push(node.(html.Node))
	}

	return body, nil
}
