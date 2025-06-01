package markdown

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/emojis"
	"github.com/thegogod/mde/html"
)

type Parser struct {
	iter *Iterator
}

func New() *Parser {
	return &Parser{
		iter: &Iterator{},
	}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	self.iter.Reset(src)

	if !self.iter.Next() {
		return nil, nil
	}

	group := html.Fragment()

	for {
		if self.iter.Curr.Kind == Eof {
			break
		}

		node, err := self.parseBlock()

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

func (self *Parser) parseBlock() (core.Node, error) {
	if self.iter.Match(Eof) {
		return nil, nil
	}

	var node core.Node = nil
	var err error = nil

	self.iter.Save()

	for range self.iter.blockQuoteDepth - 1 {
		if !self.iter.Match(BlockQuote) {
			break
		}
	}

	if self.iter.Match(H1) {
		node, err = self.parseHeading(1)
	} else if self.iter.Match(H2) {
		node, err = self.parseHeading(2)
	} else if self.iter.Match(H3) {
		node, err = self.parseHeading(3)
	} else if self.iter.Match(H4) {
		node, err = self.parseHeading(4)
	} else if self.iter.Match(H5) {
		node, err = self.parseHeading(5)
	} else if self.iter.Match(H6) {
		node, err = self.parseHeading(6)
	} else if self.iter.Match(Hr) {
		node, err = self.parseHr()
	} else if self.iter.Match(CodeBlock) {
		node, err = self.parseCodeBlock()
	} else if self.iter.Match(BlockQuote) {
		node, err = self.parseBlockQuote()
	} else if self.iter.Match(Ul) {
		node, err = self.parseUnorderedList()
	} else if self.iter.Match(Ol) {
		node, err = self.parseOrderedList()
	} else if self.iter.Match(NewLine) {
		node, err = self.parseBlock()
	}

	if err != nil {
		self.iter.Revert()
	}

	if node == nil || err != nil {
		node, err = self.parseParagraph()
	}

	if err != nil {
		self.iter.Revert()
	}

	self.iter.Pop()
	return node, err
}

func (self *Parser) parseInline() (core.Node, error) {
	if self.iter.Match(Eof) {
		return nil, nil
	}

	var node core.Node = nil
	var err error = nil

	self.iter.Save()

	if self.iter.Match(Bold) {
		node, err = self.parseBold()
	} else if self.iter.Match(BoldAlt) {
		node, err = self.parseBoldAlt()
	} else if self.iter.Match(Italic) {
		node, err = self.parseItalic()
	} else if self.iter.Match(ItalicAlt) {
		node, err = self.parseItalicAlt()
	} else if self.iter.Match(Strike) {
		node, err = self.parseStrike()
	} else if self.iter.Match(StrikeAlt) {
		node, err = self.parseStrikeAlt()
	} else if self.iter.Match(Br) {
		node, err = self.parseBr()
	} else if self.iter.Match(Code) {
		node, err = self.parseCode()
	} else if self.iter.Match(LeftBracket) {
		node, err = self.parseLink()
	} else if self.iter.Match(Bang) {
		node, err = self.parseImage()
	} else if self.iter.Match(Colon) {
		node, err = self.parseEmoji()
	} else if self.iter.Match(NewLine) {
		if self.iter.Match(NewLine) {
			self.iter.Pop()
			return nil, nil
		}

		for range self.iter.blockQuoteDepth {
			if !self.iter.Match(BlockQuote) {
				break
			}
		}

		node, err = self.parseNewLine()
	}

	if err != nil {
		self.iter.Revert()
	}

	if node == nil || err != nil {
		text, texterr := self.parseText()

		if text == nil || texterr != nil {
			return text, texterr
		}

		for self.iter.Curr.Kind == Text {
			node, err := self.parseText()

			if node == nil || err != nil {
				return text, err
			}

			text = append(text, node...)
		}

		node = text
		err = nil
	}

	self.iter.Pop()
	return node, err
}

//
// Blocks
//

func (self *Parser) parseHeading(depth int) (core.Node, error) {
	heading := html.Heading(depth)

	for self.iter.Curr.Kind.IsInline() {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}

func (self *Parser) parseParagraph() (core.Node, error) {
	paragraph := html.P()
	buff := html.Raw{}

	for self.iter.Curr.Kind.IsInline() {
		node, err := self.parseInline()

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

func (self *Parser) parseCodeBlock() (core.Node, error) {
	code := html.Code()
	lang, err := self.parseTextUntil(NewLine)

	if lang == nil || err != nil {
		return html.Pre(code), err
	}

	if len(lang) > 0 {
		code.Class(fmt.Sprintf("language-%s", lang))
	}

	node, err := self.parseTextUntil(CodeBlock)

	if node == nil || err != nil {
		return html.Pre(code), err
	}

	code.Push(node[:len(node)-1])
	return html.Pre(code), nil
}

func (self *Parser) parseBlockQuote() (core.Node, error) {
	self.iter.blockQuoteDepth++
	blockQuote := html.BlockQuote()

	for {
		node, err := self.parseBlock()

		if node == nil || err != nil {
			self.iter.blockQuoteDepth--
			return blockQuote, err
		}

		blockQuote.Push(node)

		if self.iter.Curr.Kind != BlockQuote {
			break
		}
	}

	self.iter.blockQuoteDepth--
	return blockQuote, nil
}

func (self *Parser) parseUnorderedList() (core.Node, error) {
	self.iter.listDepth++
	ul := html.Ul()

	for {
		node, err := self.parseListItem()

		if node == nil || err != nil {
			self.iter.listDepth--
			return ul, err
		}

		ul.Push(node)

		if !self.iter.MatchCount(Tab, self.iter.listDepth-1) {
			break
		}

		if !self.iter.Match(Ul) {
			break
		}
	}

	self.iter.listDepth--
	return ul, nil
}

func (self *Parser) parseOrderedList() (core.Node, error) {
	self.iter.listDepth++
	ol := html.Ol()

	for {
		node, err := self.parseListItem()

		if node == nil || err != nil {
			self.iter.listDepth--
			return ol, err
		}

		ol.Push(node)

		if !self.iter.MatchCount(Tab, self.iter.listDepth-1) {
			break
		}

		if !self.iter.Match(Ol) {
			break
		}
	}

	self.iter.listDepth--
	return ol, nil
}

//
// InLines
//

func (self *Parser) parseBold() (core.Node, error) {
	bold := html.Strong()

	for !self.iter.Match(Bold) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}

func (self *Parser) parseBoldAlt() (core.Node, error) {
	bold := html.Strong()

	for !self.iter.Match(BoldAlt) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return bold, err
		}

		bold.Push(node)
	}

	return bold, nil
}

func (self *Parser) parseItalic() (core.Node, error) {
	italic := html.I()

	for !self.iter.Match(Italic) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}

func (self *Parser) parseItalicAlt() (core.Node, error) {
	italic := html.I()

	for !self.iter.Match(ItalicAlt) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}

func (self *Parser) parseStrike() (core.Node, error) {
	strike := html.S()

	for !self.iter.Match(Strike) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}

func (self *Parser) parseStrikeAlt() (core.Node, error) {
	strike := html.S()

	for !self.iter.Match(StrikeAlt) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}

func (self *Parser) parseBr() (core.Node, error) {
	return html.Br(), nil
}

func (self *Parser) parseCode() (core.Node, error) {
	code := html.Code()
	text, err := self.parseTextUntil(Code)

	if text == nil || err != nil {
		return code, err
	}

	code.Push(text)
	return code, nil
}

func (self *Parser) parseLink() (core.Node, error) {
	link := html.A()

	for !self.iter.Match(RightBracket) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return link, err
		}

		link.Push(node)
	}

	if _, err := self.iter.Consume(LeftParen, "expected '('"); err != nil {
		return link, err
	}

	node, err := self.parseTextUntil(RightParen)

	if node == nil || err != nil {
		return link, err
	}

	link.Href(node.String())
	return link, nil
}

func (self *Parser) parseImage() (core.Node, error) {
	image := html.Img()

	if _, err := self.iter.Consume(LeftBracket, "expected '['"); err != nil {
		return image, err
	}

	node, err := self.parseTextUntil(RightBracket)

	if node == nil || err != nil {
		return image, err
	}

	image.Alt(node.String())

	if _, err := self.iter.Consume(LeftParen, "expected '('"); err != nil {
		return image, err
	}

	node, err = self.parseTextUntil(RightParen)

	if node == nil || err != nil {
		return image, err
	}

	image.Src(node.String())
	return image, nil
}

func (self *Parser) parseListItem() (*html.ListItemElement, error) {
	li := html.Li()
	self.iter.Save()
	node, err := self.parseTask()

	if err == nil && node != nil {
		li.Push(node)
		self.iter.Pop()
		return li, nil
	}

	self.iter.Revert()

	for self.iter.Curr.Kind.IsInline() {
		node, err := self.parseInline()

		if err != nil {
			self.iter.Revert()
			self.iter.Pop()
			return li, err
		}

		if node == nil {
			break
		}

		if node.String() == "\n" {
			if !self.iter.MatchCount(Tab, self.iter.listDepth) {
				break
			}

			node, err = nil, nil
			self.iter.Save()

			if self.iter.Match(Ol) {
				node, err = self.parseOrderedList()
			} else if self.iter.Match(Ul) {
				node, err = self.parseUnorderedList()
			}

			if node != nil && err == nil {
				li.Push(node)
			} else {
				self.iter.Revert()
			}

			self.iter.Pop()
			break
		}

		li.Push(node)
	}

	self.iter.Pop()
	return li, nil
}

func (self *Parser) parseTask() (core.Node, error) {
	id := uuid.NewString()
	label := html.Label().For(id)
	input := html.CheckBoxInput().Id(id)
	_, err := self.iter.Consume(LeftBracket, "expected '['")

	if err != nil {
		return label, err
	}

	checked, err := self.iter.Consume(Text, "expected ' ' or 'x'")

	if err != nil {
		return label, err
	}

	if checked.String() != " " && checked.String() != "x" {
		return label, errors.New("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		input.Checked(true)
	}

	_, err = self.iter.Consume(RightBracket, "expected ']'")

	if err != nil {
		return label, err
	}

	space, err := self.iter.Consume(Text, "expected ' '")

	if err != nil {
		return label, err
	}

	if space.String() != " " {
		return label, errors.New("expected ' '")
	}

	text := ""

	for !self.iter.Match(NewLine) {
		node, err := self.parseText()

		if err != nil {
			return label, err
		}

		if node == nil {
			break
		}

		text += node.String()
	}

	label.Push(input)
	label.Push(html.Span(text))
	return label, nil
}

func (self *Parser) parseEmoji() (html.Raw, error) {
	alias := html.Raw{}

	for !self.iter.Match(Colon) {
		node, err := self.parseText()

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

func (self *Parser) parseNewLine() (html.Raw, error) {
	if self.iter.Curr.IsWhitespace() {
		return nil, nil
	}

	return html.Raw("\n"), nil
}

func (self *Parser) parseText() (html.Raw, error) {
	if self.iter.Curr.Kind == Eof {
		return nil, nil
	}

	text := html.Raw(self.iter.Curr.Value)
	self.iter.Next()
	return text, nil
}

func (self *Parser) parseTextUntil(kind TokenKind) (html.Raw, error) {
	if self.iter.Curr.Kind == Eof {
		return nil, nil
	}

	text := html.Raw{}

	for !self.iter.Match(kind) {
		node, err := self.parseText()

		if node == nil || err != nil {
			return text, err
		}

		text = append(text, node...)
	}

	return text, nil
}
