package syntax

import (
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

func (self CodeBlock) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.MatchCount(core.BackQuote, 3)
}

func (self CodeBlock) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	code := html.Code()
	lang, err := parser.ParseTextUntil(core.NewLine, parser, iter)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.Class(fmt.Sprintf("language-%s", lang))
	}

	buff := html.Raw{}

	for !iter.MatchCount(core.BackQuote, 3) {
		node, err := parser.ParseText(parser, iter)

		if node == nil {
			return html.Pre(code), iter.Curr().Error("expected closing '```'")
		}

		if err != nil {
			return html.Pre(code), err
		}

		if string(node) == "\n" {
			buff = append(buff, node...)
			continue
		}

		if len(buff) > 0 {
			code.Push(buff)
			buff = html.Raw{}
		}

		code.Push(node)
	}

	return html.Pre(code), nil
}
