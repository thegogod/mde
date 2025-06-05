package syntax

import (
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type CodeBlock struct{}

func (self CodeBlock) IsBlock() bool {
	return true
}

func (self CodeBlock) IsInline() bool {
	return false
}

func (self CodeBlock) Name() string {
	return "code_block"
}

func (self CodeBlock) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.Match(tokens.CodeBlock)
}

func (self CodeBlock) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	code := html.Code()
	lang, err := parser.ParseTextUntil(iter, tokens.NewLine)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.Class(fmt.Sprintf("language-%s", lang))
	}

	node, err := parser.ParseTextUntil(iter, tokens.CodeBlock)

	if node == nil {
		return html.Pre(code), iter.Curr().Error("expected closing '```'")
	}

	if err != nil {
		return html.Pre(code), err
	}

	code.Push(node[:len(node)-1])
	return html.Pre(code), nil
}
