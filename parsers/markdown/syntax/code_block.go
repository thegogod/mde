package syntax

import (
	"errors"
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type CodeBlock struct{}

func (self *CodeBlock) IsBlock() bool {
	return true
}

func (self *CodeBlock) IsInline() bool {
	return false
}

func (self *CodeBlock) Select(iter core.Iterator) bool {
	return iter.Curr().Kind() == tokens.CodeBlock
}

func (self *CodeBlock) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	code := html.Code()

	if _, err := iter.Consume(tokens.CodeBlock, "expected '```'"); err != nil {
		return html.Pre(code), err
	}

	lang, err := parser.ParseTextUntil(iter, tokens.NewLine)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.Class(fmt.Sprintf("language-%s", lang))
	}

	node, err := parser.ParseTextUntil(iter, tokens.CodeBlock)

	if node == nil {
		return html.Pre(code), errors.New("expected closing '```'")
	}

	if err != nil {
		return html.Pre(code), err
	}

	code.Push(node[:len(node)-1])
	return html.Pre(code), nil
}
