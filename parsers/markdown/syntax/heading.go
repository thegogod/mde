package syntax

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Heading struct{}

func (self *Heading) IsBlock() bool {
	return true
}

func (self *Heading) IsInline() bool {
	return false
}

func (self *Heading) Select(iter core.Iterator) bool {
	switch iter.Curr().Kind() {
	case tokens.H1:
		fallthrough
	case tokens.H2:
		fallthrough
	case tokens.H3:
		fallthrough
	case tokens.H4:
		fallthrough
	case tokens.H5:
		fallthrough
	case tokens.H6:
		return true
	default:
		return false
	}
}

func (self *Heading) Parse(parser core.Parser, iter core.Iterator) (core.Node, error) {
	depth := len(iter.Curr().Bytes())
	heading := html.Heading(depth)

	if _, err := iter.Consume(byte(9+depth), "expected '#' or '##' or '###' or '####' or '#####' or '######'"); err != nil {
		return heading, err
	}

	for iter.Curr().IsInline() {
		node, err := parser.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
