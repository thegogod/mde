package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/emojis"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/tokens"
)

type Emoji struct{}

func (self Emoji) IsBlock() bool {
	return false
}

func (self Emoji) IsInline() bool {
	return true
}

func (self Emoji) Name() string {
	return "emoji"
}

func (self Emoji) Select(parser core.Parser, iter core.Iterator) bool {
	return iter.MatchCount(tokens.Colon, 1)
}

func (self Emoji) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	alias := html.Raw{}

	for !iter.Match(tokens.Colon) {
		node, err := parser.ParseText(iter)

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
