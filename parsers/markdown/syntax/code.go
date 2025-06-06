package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Code struct{}

func (self Code) IsBlock() bool {
	return false
}

func (self Code) IsInline() bool {
	return true
}

func (self Code) Name() string {
	return "code"
}

func (self Code) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.MatchCount(tokens.BackQuote, 1)
}

func (self Code) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	code := html.Code()
	text, err := parser.ParseTextUntil(iter, tokens.BackQuote)

	if text == nil {
		return code, iter.Curr().Error("expected closing '`'")
	}

	if err != nil {
		return code, err
	}

	code.Push(text)
	return code, nil
}
