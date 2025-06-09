package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Paragraph struct{}

func (self Paragraph) IsBlock() bool {
	return true
}

func (self Paragraph) IsInline() bool {
	return false
}

func (self Paragraph) Name() string {
	return "paragraph"
}

func (self Paragraph) Select(parser core.Parser, iter *core.Iterator) bool {
	return true
}

func (self Paragraph) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	paragraph := html.P()
	buff := html.Raw{}

	for iter.Curr().Kind() != tokens.Eof {
		node, err := parser.ParseInline(iter)

		if node == nil {
			break
		}

		if err != nil {
			return paragraph, err
		}

		if node.String() == "\n" {
			if iter.Curr().Kind() == tokens.GreaterThan {
				break
			}

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
