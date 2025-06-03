package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func (self *Parser) parseParagraph(iter *tokens.Iterator) (core.Node, error) {
	paragraph := html.P()
	buff := html.Raw{}

	for iter.Curr().IsInline() {
		node, err := self.ParseInline(iter)

		if err != nil {
			return paragraph, err
		}

		if node == nil {
			break
		}

		if node.String() == "\n" {
			buff = append(buff, node.Bytes()...)
			continue
		}

		if len(buff) > 0 {
			paragraph.Push(buff)
			buff = html.Raw{}
		}

		paragraph.Push(node)
	}

	if len(paragraph.Children()) == 0 {
		return nil, nil
	}

	return paragraph, nil
}
