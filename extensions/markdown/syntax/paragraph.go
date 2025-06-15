package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
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

	for iter.Curr().Kind() != core.Eof {
		node, err := parser.ParseInline(parser, iter)

		if node == nil {
			break
		}

		if err != nil {
			return paragraph, err
		}

		if string(node.Render()) == "\n" {
			if iter.Curr().Kind() == core.GreaterThan {
				break
			}

			buff = append(buff, node.Render()...)
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
