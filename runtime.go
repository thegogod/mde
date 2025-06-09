package mde

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Runtime struct {
	_parsers []core.Parser
}

func New(parsers ...core.Parser) *Runtime {
	return &Runtime{
		_parsers: append(parsers, markdown.New()),
	}
}

func (self Runtime) Parse(src []byte) (core.Node, error) {
	iter := core.Iter(tokens.NewScanner(src))

	if !iter.Next() {
		return nil, nil
	}

	group := html.Fragment()

	for {
		if iter.Curr().Kind() == tokens.Eof {
			break
		}

		var node core.Node
		var err error

		for _, parser := range self._parsers {
			node, err = parser.ParseBlock(parser, iter)

			if err != nil {
				return group, err
			}

			if node == nil {
				continue
			}
		}

		group.Push(node)
	}

	return group, nil
}
