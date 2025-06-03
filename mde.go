package mde

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Mde struct {
	parsers []core.Parser
}

func New(parsers ...core.Parser) *Mde {
	return &Mde{
		parsers: append(parsers, markdown.New()),
	}
}

func (self *Mde) Parse(src []byte) (core.Node, error) {
	iter := &tokens.Iterator{}
	iter.Reset(src)

	if !iter.Next() {
		return nil, nil
	}

	group := html.Fragment()

	for {
		var node core.Node
		var err error

		if iter.Curr().Kind() == tokens.Eof {
			break
		}

		for _, parser := range self.parsers {
			node, err = parser.ParseBlock(iter)

			if err == nil {
				break
			}
		}

		if node == nil && err == nil {
			continue
		}

		if err != nil {
			return nil, err
		}

		group.Push(node)
	}

	return group, nil
}
