package markdown

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/emojis"
	"github.com/thegogod/mde/html"
	"github.com/thegogod/mde/parsers/markdown/tokens"
)

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	iter := &Iterator{}
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

	iter := iterator.(*Iterator)
	var node core.Node = nil
	var err error = nil

	iter.Save()

	for range iter.blockQuoteDepth - 1 {
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

	iter := iterator.(*Iterator)
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

		for range iter.blockQuoteDepth {
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

//
// Blocks
//

func (self *Parser) parseHeading(depth int, iter *Iterator) (core.Node, error) {
	heading := html.Heading(depth)

	for iter.Curr().IsInline() {
		node, err := self.ParseInline(iter)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}

func (self *Parser) parseParagraph(iter *Iterator) (core.Node, error) {
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

func (self *Parser) parseHr() (core.Node, error) {
	return html.Hr(), nil
}

func (self *Parser) parseCodeBlock(iter *Iterator) (core.Node, error) {
	code := html.Code()
	lang, err := self.ParseTextUntil(iter, tokens.NewLine)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.Class(fmt.Sprintf("language-%s", lang))
	}

	node, err := self.ParseTextUntil(iter, tokens.CodeBlock)

	if node == nil {
		return html.Pre(code), errors.New("expected closing '```'")
	}

	if err != nil {
		return html.Pre(code), err
	}

	code.Push(node[:len(node)-1])
	return html.Pre(code), nil
}

func (self *Parser) parseBlockQuote(iter *Iterator) (core.Node, error) {
	iter.blockQuoteDepth++
	blockQuote := html.BlockQuote()

	for {
		node, err := self.ParseBlock(iter)

		if node == nil || err != nil {
			iter.blockQuoteDepth--
			return blockQuote, err
		}

		blockQuote.Push(node)

		if iter.Curr().Kind() != tokens.BlockQuote {
			break
		}
	}

	iter.blockQuoteDepth--
	return blockQuote, nil
}

func (self *Parser) parseUnorderedList(iter *Iterator) (core.Node, error) {
	iter.listDepth++
	ul := html.Ul()

	for {
		node, err := self.parseListItem(iter)

		if node == nil || err != nil {
			iter.listDepth--
			return ul, err
		}

		ul.Push(node)

		if !iter.MatchCount(tokens.Tab, iter.listDepth-1) {
			break
		}

		if !iter.Match(tokens.Ul) {
			break
		}
	}

	iter.listDepth--
	return ul, nil
}

func (self *Parser) parseOrderedList(iter *Iterator) (core.Node, error) {
	iter.listDepth++
	ol := html.Ol()

	for {
		node, err := self.parseListItem(iter)

		if node == nil || err != nil {
			iter.listDepth--
			return ol, err
		}

		ol.Push(node)

		if !iter.MatchCount(tokens.Tab, iter.listDepth-1) {
			break
		}

		if !iter.Match(tokens.Ol) {
			break
		}
	}

	iter.listDepth--
	return ol, nil
}

//
// InLines
//

func (self *Parser) parseBold(iter *Iterator) (core.Node, error) {
	bold := html.Strong()

	for !iter.Match(tokens.Bold) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return bold, errors.New("expected closing '**'")
		}

		if err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}

func (self *Parser) parseBoldAlt(iter *Iterator) (core.Node, error) {
	bold := html.Strong()

	for !iter.Match(tokens.BoldAlt) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return bold, errors.New("expected closing '__'")
		}

		if err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}

func (self *Parser) parseItalic(iter *Iterator) (core.Node, error) {
	italic := html.I()

	for !iter.Match(tokens.Italic) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return italic, errors.New("expected closing '*'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}

func (self *Parser) parseItalicAlt(iter *Iterator) (core.Node, error) {
	italic := html.I()

	for !iter.Match(tokens.ItalicAlt) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return italic, errors.New("expected closing '_'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}

func (self *Parser) parseStrike(iter *Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(tokens.Strike) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return strike, errors.New("expected closing '~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}

func (self *Parser) parseStrikeAlt(iter *Iterator) (core.Node, error) {
	strike := html.S()

	for !iter.Match(tokens.StrikeAlt) {
		node, err := self.ParseInline(iter)

		if node == nil {
			return strike, errors.New("expected closing '~~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}

func (self *Parser) parseBr() (core.Node, error) {
	return html.Br(), nil
}

func (self *Parser) parseCode(iter *Iterator) (core.Node, error) {
	code := html.Code()
	text, err := self.ParseTextUntil(iter, tokens.Code)

	if text == nil {
		return code, errors.New("expected closing '`'")
	}

	if err != nil {
		return code, err
	}

	code.Push(text)
	return code, nil
}

func (self *Parser) parseLink(iter *Iterator) (core.Node, error) {
	link := html.A()

	for !iter.Match(tokens.RightBracket) {
		node, err := self.ParseInline(iter)

		if node == nil || err != nil {
			return link, err
		}

		link.Push(node)
	}

	if _, err := iter.Consume(tokens.LeftParen, "expected '('"); err != nil {
		return link, err
	}

	node, err := self.ParseTextUntil(iter, tokens.RightParen)

	if node == nil || err != nil {
		return link, err
	}

	link.Href(string(node))
	return link, nil
}

func (self *Parser) parseImage(iter *Iterator) (core.Node, error) {
	image := html.Img()

	if _, err := iter.Consume(tokens.LeftBracket, "expected '['"); err != nil {
		return image, err
	}

	node, err := self.ParseTextUntil(iter, tokens.RightBracket)

	if node == nil || err != nil {
		return image, err
	}

	image.Alt(string(node))

	if _, err := iter.Consume(tokens.LeftParen, "expected '('"); err != nil {
		return image, err
	}

	node, err = self.ParseTextUntil(iter, tokens.RightParen)

	if node == nil || err != nil {
		return image, err
	}

	image.Src(string(node))
	return image, nil
}

func (self *Parser) parseListItem(iter *Iterator) (*html.ListItemElement, error) {
	li := html.Li()
	iter.Save()
	node, err := self.parseTask(iter)

	if err == nil && node != nil {
		li.Push(node)
		iter.Pop()
		return li, nil
	}

	iter.Revert()

	for iter.Curr().IsInline() {
		node, err := self.ParseInline(iter)

		if err != nil {
			iter.Revert()
			iter.Pop()
			return li, err
		}

		if node == nil {
			break
		}

		if node.String() == "\n" {
			if !iter.MatchCount(tokens.Tab, iter.listDepth) {
				break
			}

			node, err = nil, nil
			iter.Save()

			if iter.Match(tokens.Ol) {
				node, err = self.parseOrderedList(iter)
			} else if iter.Match(tokens.Ul) {
				node, err = self.parseUnorderedList(iter)
			}

			if node != nil && err == nil {
				li.Push(node)
			} else {
				iter.Revert()
			}

			iter.Pop()
			break
		}

		li.Push(node)
	}

	iter.Pop()
	return li, nil
}

func (self *Parser) parseTask(iter *Iterator) (core.Node, error) {
	id := uuid.NewString()
	label := html.Label().For(id)
	input := html.CheckBoxInput().Id(id)
	_, err := iter.Consume(tokens.LeftBracket, "expected '['")

	if err != nil {
		return label, err
	}

	checked, err := iter.Consume(tokens.Text, "expected ' ' or 'x'")

	if err != nil {
		return label, err
	}

	if checked.String() != " " && checked.String() != "x" {
		return label, errors.New("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		input.Checked(true)
	}

	_, err = iter.Consume(tokens.RightBracket, "expected ']'")

	if err != nil {
		return label, err
	}

	space, err := iter.Consume(tokens.Text, "expected ' '")

	if err != nil {
		return label, err
	}

	if space.String() != " " {
		return label, errors.New("expected ' '")
	}

	text := ""

	for !iter.Match(tokens.NewLine) {
		node, err := self.ParseText(iter)

		if err != nil {
			return label, err
		}

		if node == nil {
			break
		}

		text += string(node)
	}

	label.Push(input)
	label.Push(html.Span(text))
	return label, nil
}

func (self *Parser) parseEmoji(iter *Iterator) (html.Raw, error) {
	alias := html.Raw{}

	for !iter.Match(tokens.Colon) {
		node, err := self.ParseText(iter)

		if node == nil {
			return alias, errors.New("expected closing ':'")
		}

		if err != nil {
			return alias, err
		}

		alias = append(alias, node...)
	}

	emoji, exists := emojis.Get(alias.String())

	if !exists {
		return alias, errors.New("emoji alias not found")
	}

	return html.Raw(emoji.Emoji), nil
}

func (self *Parser) parseNewLine(iter *Iterator) (html.Raw, error) {
	curr := iter.Curr().String()

	if curr == " " || curr == "\n" {
		return nil, nil
	}

	return html.Raw("\n"), nil
}
