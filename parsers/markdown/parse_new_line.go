package markdown

import (
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseNewLine(iter *tokens.Iterator) (html.Raw, error) {
	curr := iter.Curr().String()

	if curr == " " || curr == "\n" {
		return nil, nil
	}

	return html.Raw("\n"), nil
}
