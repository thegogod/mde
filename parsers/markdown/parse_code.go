package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseCode(iter *tokens.Iterator) (core.Node, error) {
	code := html.Code()
	text, err := self.parseTextUntil(iter, tokens.Code)

	if text == nil {
		return code, iter.Curr().Error("expected closing '`'")
	}

	if err != nil {
		return code, err
	}

	code.Push(text)
	return code, nil
}
