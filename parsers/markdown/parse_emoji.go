package markdown

import (
	"github.com/thegogod/mde/emojis"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseEmoji(iter *tokens.Iterator) (html.Raw, error) {
	alias := html.Raw{}

	for !iter.Match(tokens.Colon) {
		node, err := self.parseText(iter)

		if node == nil {
			return alias, iter.Curr().Error("expected closing ':'")
		}

		if err != nil {
			return alias, err
		}

		alias = append(alias, node...)
	}

	emoji, exists := emojis.Get(alias.String())

	if !exists {
		return alias, iter.Curr().Error("emoji alias not found")
	}

	return html.Raw(emoji.Emoji), nil
}
