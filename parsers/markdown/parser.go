package markdown

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/syntax"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Parser struct {
	syntax []core.Syntax
}

func New() *Parser {
	return &Parser{
		syntax: []core.Syntax{
			syntax.Heading{},
			syntax.HorizontalRule{},
			syntax.CodeBlock{},
			syntax.BlockQuote{},
			syntax.UnOrderedList{},
			syntax.OrderedList{},
			syntax.Paragraph{},

			syntax.Bold{},
			syntax.BoldAlt{},
			syntax.Italic{},
			syntax.ItalicAlt{},
			syntax.Strike{},
			syntax.StrikeAlt{},
			syntax.Highlight{},
			syntax.Code{},
			syntax.Emoji{},
			syntax.Link{},
			syntax.Url{},
			syntax.Image{},
			syntax.Break{},
			syntax.ListItem{},
			syntax.NewLine{},
		},
	}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	iter := &tokens.Iterator{}
	iter.Reset(src)

	if !iter.Next() {
		return nil, nil
	}

	group := html.Fragment()

	for {
		if iter.Curr().Kind() == tokens.Eof {
			break
		}

		node, err := self.ParseBlock(iter)

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

func (self *Parser) ParseBlock(iterator core.Iterator) (core.Node, error) {
	if iterator.Match(tokens.Eof) {
		return nil, nil
	}

	iter := iterator.(*tokens.Iterator)
	var node core.Node = nil
	var err error = nil

	iterator.Save()

	for range iter.BlockQuoteDepth - 1 {
		if !iterator.Match(tokens.BlockQuote) {
			break
		}
	}

	if iterator.Match(tokens.NewLine) {
		iterator.Pop()
		return self.ParseBlock(iterator)
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

func (self *Parser) ParseInline(iterator core.Iterator) (core.Node, error) {
	if iterator.Match(tokens.Eof) {
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
		text, texterr := self.ParseText(iterator)

		if text != nil {
			node = html.Raw(text)
		}

		err = texterr
	}

	iterator.Pop()
	return node, err
}

func (self *Parser) ParseSyntax(name string, iter core.Iterator) (core.Node, error) {
	i := slices.IndexFunc(self.syntax, func(s core.Syntax) bool {
		return s.Name() == name
	})

	if i < 0 {
		panic(fmt.Sprintf("syntax '%s' not found", name))
	}

	return self.syntax[i].Parse(self, iter)
}

func (self *Parser) ParseText(iter core.Iterator) ([]byte, error) {
	if iter.Curr().Kind() == tokens.Eof {
		return nil, nil
	}

	text := html.Raw(iter.Curr().Bytes())
	iter.Next()

	if bytes.Equal(text, []byte{' '}) {
		return text, nil
	}

	for iter.Curr().Kind() == tokens.Text {
		if bytes.Equal(iter.Curr().Bytes(), []byte{' '}) {
			return text, nil
		}

		text = append(text, iter.Curr().Bytes()...)
		iter.Next()
	}

	return text, nil
}

func (self *Parser) ParseTextUntil(iter core.Iterator, kind tokens.TokenKind) ([]byte, error) {
	if iter.Curr().Kind() == tokens.Eof {
		return nil, nil
	}

	text := html.Raw{}

	for !iter.Match(kind) {
		node, err := self.ParseText(iter)

		if node == nil || err != nil {
			return text, err
		}

		text = append(text, node...)
	}

	return text, nil
}
