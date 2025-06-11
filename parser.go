package mde

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/markdown"
	"github.com/thegogod/mde/html"
)

type Parser struct {
	_extensions core.ExtensionRegistry
}

func New(extensions ...core.Extension) *Parser {
	registry := core.ExtensionRegistry{}

	for _, ext := range extensions {
		registry.Add(ext)
	}

	if !registry.Has("markdown") {
		registry.Add(markdown.New())
	}

	return &Parser{
		_extensions: registry,
	}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	scanner := core.NewScanner(src)

	for _, ext := range self._extensions {
		scanner.Extend(ext)
	}

	iter := core.Iter(scanner)

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

	for _, ext := range self._extensions {
		for _, syntax := range ext.Syntax() {
			if syntax.IsBlock() && syntax.Select(self, iterator) {
				node, err = syntax.Parse(self, iterator)

				if err == nil {
					break
				}
			}

			iterator.Revert()
		}
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

	for _, ext := range self._extensions {
		for _, syntax := range ext.Syntax() {
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

func (self *Parser) ParseSyntax(extension string, name string, parser core.Parser, iter *core.Iterator) (core.Node, error) {
	ext := self._extensions.Get(extension)

	if ext == nil {
		panic(fmt.Sprintf("extension '%s' not found", extension))
	}

	syntax := ext.Syntax()

	i := slices.IndexFunc(syntax, func(s core.Syntax) bool {
		return s.Name() == name
	})

	if i < 0 {
		panic(fmt.Sprintf("syntax '%s' not found", name))
	}

	return syntax[i].Parse(self, iter)
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
