package markdown

import (
	"slices"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Parser struct {
	syntax []core.Syntax
}

func New() *Parser {
	return &Parser{
		syntax: []core.Syntax{},
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

func (self *Parser) ParseBlock(iterator core.Iterator) (core.Node, error) {
	if iterator.Match(tokens.Eof) {
		return nil, nil
	}

	iter := iterator.(*tokens.Iterator)
	var node core.Node = nil
	var err error = nil

	iter.Save()

	for range iter.BlockQuoteDepth - 1 {
		if !iter.Match(tokens.BlockQuote) {
			break
		}
	}

	if iter.Match(tokens.H1) {
		node, err = self.parseHeading(1, iter)
	} else if iter.Match(tokens.H2) {
		node, err = self.parseHeading(2, iter)
	} else if iter.Match(tokens.H3) {
		node, err = self.parseHeading(3, iter)
	} else if iter.Match(tokens.H4) {
		node, err = self.parseHeading(4, iter)
	} else if iter.Match(tokens.H5) {
		node, err = self.parseHeading(5, iter)
	} else if iter.Match(tokens.H6) {
		node, err = self.parseHeading(6, iter)
	} else if iter.Match(tokens.Hr) {
		node, err = self.parseHr()
	} else if iter.Match(tokens.CodeBlock) {
		node, err = self.parseCodeBlock(iter)
	} else if iter.Match(tokens.BlockQuote) {
		node, err = self.parseBlockQuote(iter)
	} else if iter.Match(tokens.Ul) {
		node, err = self.parseUnorderedList(iter)
	} else if iter.Match(tokens.Ol) {
		node, err = self.parseOrderedList(iter)
	} else if iter.Match(tokens.NewLine) {
		node, err = self.ParseBlock(iter)
	}

	if err != nil {
		iter.Revert()
	}

	if node == nil || err != nil {
		node, err = self.parseParagraph(iter)
	}

	if err != nil {
		iter.Revert()
	}

	iter.Pop()
	return node, err
}

func (self *Parser) ParseInline(iterator core.Iterator) (core.Node, error) {
	if iterator.Match(tokens.Eof) {
		return nil, nil
	}

	iter := iterator.(*tokens.Iterator)
	var node core.Node = nil
	var err error = nil

	iter.Save()

	if iter.Match(tokens.Bold) {
		node, err = self.parseBold(iter)
	} else if iter.Match(tokens.BoldAlt) {
		node, err = self.parseBoldAlt(iter)
	} else if iter.Match(tokens.Italic) {
		node, err = self.parseItalic(iter)
	} else if iter.Match(tokens.ItalicAlt) {
		node, err = self.parseItalicAlt(iter)
	} else if iter.Match(tokens.Strike) {
		node, err = self.parseStrike(iter)
	} else if iter.Match(tokens.StrikeAlt) {
		node, err = self.parseStrikeAlt(iter)
	} else if iter.Match(tokens.Br) {
		node, err = self.parseBr()
	} else if iter.Match(tokens.Code) {
		node, err = self.parseCode(iter)
	} else if iter.Match(tokens.LeftBracket) {
		node, err = self.parseLink(iter)
	} else if iter.Match(tokens.Bang) {
		node, err = self.parseImage(iter)
	} else if iter.Match(tokens.Colon) {
		node, err = self.parseEmoji(iter)
	} else if iter.Match(tokens.NewLine) {
		if iter.Match(tokens.NewLine) {
			iter.Pop()
			return nil, nil
		}

		for range iter.BlockQuoteDepth {
			if !iter.Match(tokens.BlockQuote) {
				break
			}
		}

		node, err = self.parseNewLine(iter)
	}

	if err != nil {
		iter.Revert()
	}

	if node == nil || err != nil {
		text, texterr := self.ParseText(iter)

		if text == nil || texterr != nil {
			return html.Raw(text), texterr
		}

		for iter.Curr().Kind() == tokens.Text {
			node, err := self.ParseText(iter)

			if node == nil || err != nil {
				return html.Raw(text), err
			}

			text = append(text, node...)
		}

		node = html.Raw(text)
		err = nil
	}

	iter.Pop()
	return node, err
}

func (self *Parser) ParseSyntax(name string, iter core.Iterator) (core.Node, error) {
	i := slices.IndexFunc(self.syntax, func(s core.Syntax) bool {
		return s.Name() == name
	})

	if i < 0 {
		return nil, nil
	}

	return self.syntax[i].Parse(self, iter)
}

func (self *Parser) ParseText(iter core.Iterator) ([]byte, error) {
	if iter.Curr().Kind() == tokens.Eof {
		return nil, nil
	}

	text := html.Raw(iter.Curr().Bytes())
	iter.Next()
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
