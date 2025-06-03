package markdown

import (
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseCodeBlock(iter *tokens.Iterator) (core.Node, error) {
	code := html.Code()
	lang, err := self.parseTextUntil(iter, tokens.NewLine)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.Class(fmt.Sprintf("language-%s", lang))
	}

	node, err := self.parseTextUntil(iter, tokens.CodeBlock)

	if node == nil {
		return html.Pre(code), iter.Curr().Error("expected closing '```'")
	}

	if err != nil {
		return html.Pre(code), err
	}

	code.Push(node[:len(node)-1])
	return html.Pre(code), nil
}
