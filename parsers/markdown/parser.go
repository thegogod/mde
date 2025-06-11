package markdown

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/syntax"
)

type Parser struct {
	syntax []core.Syntax
}

func New(rules ...core.Syntax) *Parser {
	return &Parser{
		syntax: append(rules, syntax.All...),
	}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	iter := core.Iter(core.NewScanner(src).WithTypes(Markdown{}.TokenTypes()...))

	if !iter.Next() {
		return nil, nil
	}

	group := html.Fragment()

	for {
		if iter.Curr().Kind() == core.Eof {
			break
		}

		node, err := self.ParseBlock(self, iter)

		if err != nil {
			return group, err
		}

		if node == nil {
			continue
		}

		group.Push(node)
	}

	return group, nil
}

func (self *Parser) ParseBlock(parser core.Parser, iterator *core.Iterator) (core.Node, error) {
	if iterator.Match(core.Eof) {
		return nil, nil
	}

	var node core.Node = nil
	var err error = nil

	iterator.Save()

	for range iterator.BlockQuoteDepth - 1 {
		if !iterator.Match(core.GreaterThan) {
			break
		}
	}

	if iterator.Match(core.NewLine) {
		iterator.Pop()
		return self.ParseBlock(self, iterator)
	}

	iterator.Save()

	for _, syntax := range self.syntax {
		if syntax.IsBlock() && syntax.Select(self, iterator) {
			node, err = syntax.Parse(self, iterator)

			if err == nil {
				break
			}
		}

		iterator.Revert()
	}

	iterator.Pop()
	iterator.Pop()
	return node, err
}

func (self *Parser) ParseInline(parser core.Parser, iterator *core.Iterator) (core.Node, error) {
	if iterator.Match(core.Eof) {
		return nil, nil
	}

	var node core.Node = nil
	var err error = nil

	iterator.Save()

	for _, syntax := range self.syntax {
		if syntax.IsInline() && syntax.Select(self, iterator) {
			node, err = syntax.Parse(self, iterator)

			if err == nil {
				if node == nil {
					iterator.Pop()
					return nil, nil
				}

				break
			}
		}

		iterator.Revert()
	}

	if node == nil || err != nil {
		text, texterr := self.ParseText(self, iterator)

		if text != nil {
			node = html.Raw(text)
		}

		err = texterr
	}

	iterator.Pop()
	return node, err
}

func (self *Parser) ParseSyntax(name string, parser core.Parser, iter *core.Iterator) (core.Node, error) {
	i := slices.IndexFunc(self.syntax, func(s core.Syntax) bool {
		return s.Name() == name
	})

	if i < 0 {
		panic(fmt.Sprintf("syntax '%s' not found", name))
	}

	return self.syntax[i].Parse(self, iter)
}

func (self *Parser) ParseText(parser core.Parser, iter *core.Iterator) ([]byte, error) {
	if iter.Curr().Kind() == core.Eof {
		return nil, nil
	}

	text := html.Raw(iter.Curr().Bytes())
	iter.Next()

	if bytes.Equal(text, []byte{' '}) {
		return text, nil
	}

	for iter.Curr().Kind() == core.Text {
		if bytes.Equal(iter.Curr().Bytes(), []byte{' '}) {
			return text, nil
		}

		text = append(text, iter.Curr().Bytes()...)
		iter.Next()
	}

	return text, nil
}

func (self *Parser) ParseTextUntil(kind core.TokenKind, parser core.Parser, iter *core.Iterator) ([]byte, error) {
	if iter.Curr().Kind() == core.Eof {
		return nil, nil
	}

	text := html.Raw{}

	for !iter.Match(kind) {
		node, err := self.ParseText(parser, iter)

		if node == nil || err != nil {
			return text, err
		}

		text = append(text, node...)
	}

	return text, nil
}
