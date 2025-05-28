package markdown

import (
	"errors"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/parsers/markdown/ast"
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

	group := ast.Group{
		Items: []core.Node{},
	}

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

		group.Add(node)
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

	for range self.iter.blockQuoteDepth {
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
		node, err = self.parseText()
	}

	self.iter.Pop()
	return node, err
}

//
// Blocks
//

func (self *Parser) parseHeading(depth int) (core.Node, error) {
	heading := ast.Heading{
		Depth:   depth,
		Content: []core.Node{},
	}

	for self.iter.Curr.Kind.IsInline() {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return heading, err
		}

		heading.Add(node)
	}

	return heading, nil
}

func (self *Parser) parseParagraph() (core.Node, error) {
	paragraph := ast.Paragraph{
		Content: []core.Node{},
	}

	for self.iter.Curr.Kind.IsInline() {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return paragraph, err
		}

		paragraph.Add(node)
	}

	if len(paragraph.Content) == 0 {
		return nil, nil
	}

	return paragraph, nil
}

func (self *Parser) parseHr() (core.Node, error) {
	return ast.Hr{}, nil
}

func (self *Parser) parseCodeBlock() (core.Node, error) {
	code := ast.CodeBlock{
		Content: []core.Node{},
	}

	lang := ast.Group{
		Items: []core.Node{},
	}

	for !self.iter.Match(NewLine) {
		node, _ := self.parseText()
		lang.Add(node)
	}

	if len(lang.Items) > 0 {
		code.Lang = lang
	}

	for !self.iter.Match(CodeBlock) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return code, err
		}

		code.Add(node)
	}

	return code, nil
}

func (self *Parser) parseBlockQuote() (core.Node, error) {
	self.iter.blockQuoteDepth++
	blockQuote := ast.BlockQuote{
		Content: []core.Node{},
	}

	for {
		node, err := self.parseBlock()

		if node == nil || err != nil {
			return blockQuote, err
		}

		blockQuote.Add(node)

		if self.iter.Curr.Kind != BlockQuote {
			break
		}
	}

	self.iter.blockQuoteDepth--
	return blockQuote, nil
}

func (self *Parser) parseUnorderedList() (core.Node, error) {
	ul := ast.Ul{
		Content: []ast.Li{},
	}

	for {
		node, err := self.parseListItem()

		if node == nil || err != nil {
			return ul, err
		}

		ul.Add(node.(ast.Li))

		if !self.iter.Match(Ul) {
			break
		}
	}

	return ul, nil
}

func (self *Parser) parseOrderedList() (core.Node, error) {
	ol := ast.Ol{
		Content: []ast.Li{},
	}

	for {
		node, err := self.parseListItem()

		if node == nil || err != nil {
			return ol, err
		}

		ol.Add(node.(ast.Li))

		if !self.iter.Match(Ol) {
			break
		}
	}

	return ol, nil
}

//
// InLines
//

func (self *Parser) parseBold() (core.Node, error) {
	bold := ast.Bold{
		Content: []core.Node{},
	}

	for !self.iter.Match(Bold) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return bold, err
		}

		bold.Add(node)
	}

	return bold, nil
}

func (self *Parser) parseBoldAlt() (core.Node, error) {
	bold := ast.Bold{
		Content: []core.Node{},
	}

	for !self.iter.Match(BoldAlt) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return bold, err
		}

		bold.Add(node)
	}

	return bold, nil
}

func (self *Parser) parseItalic() (core.Node, error) {
	italic := ast.Italic{
		Content: []core.Node{},
	}

	for !self.iter.Match(Italic) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return italic, err
		}

		italic.Add(node)
	}

	return italic, nil
}

func (self *Parser) parseItalicAlt() (core.Node, error) {
	italic := ast.Italic{
		Content: []core.Node{},
	}

	for !self.iter.Match(ItalicAlt) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return italic, err
		}

		italic.Add(node)
	}

	return italic, nil
}

func (self *Parser) parseStrike() (core.Node, error) {
	strike := ast.Strike{
		Content: []core.Node{},
	}

	for !self.iter.Match(Strike) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return strike, err
		}

		strike.Add(node)
	}

	return strike, nil
}

func (self *Parser) parseStrikeAlt() (core.Node, error) {
	strike := ast.Strike{
		Content: []core.Node{},
	}

	for !self.iter.Match(StrikeAlt) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return strike, err
		}

		strike.Add(node)
	}

	return strike, nil
}

func (self *Parser) parseBr() (core.Node, error) {
	return ast.Br{}, nil
}

func (self *Parser) parseCode() (core.Node, error) {
	code := ast.Code{
		Content: []core.Node{},
	}

	for !self.iter.Match(Code) {
		node, err := self.parseText()

		if node == nil || err != nil {
			return code, err
		}

		code.Add(node)
	}

	return code, nil
}

func (self *Parser) parseLink() (core.Node, error) {
	link := ast.Link{
		Text: []core.Node{},
		Href: []core.Node{},
	}

	for self.iter.Curr.Kind.IsInline() && self.iter.Curr.Kind != RightBracket {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return link, err
		}

		link.Text = append(link.Text, node)
	}

	if _, err := self.iter.Consume(RightBracket, "expected ']'"); err != nil {
		return link, err
	}

	if _, err := self.iter.Consume(LeftParen, "expected '('"); err != nil {
		return link, err
	}

	for self.iter.Curr.Kind.IsInline() && self.iter.Curr.Kind != RightParen {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return link, err
		}

		link.Href = append(link.Href, node)
	}

	if _, err := self.iter.Consume(RightParen, "expected ')'"); err != nil {
		return link, err
	}

	return link, nil
}

func (self *Parser) parseImage() (core.Node, error) {
	image := ast.Image{Alt: []core.Node{}, Src: []core.Node{}}

	if !self.iter.Match(LeftBracket) {
		return nil, nil
	}

	node, err := self.parseLink()

	if node == nil || err != nil {
		return image, err
	}

	link := node.(ast.Link)
	image.Alt = link.Text
	image.Src = link.Href
	return image, nil
}

func (self *Parser) parseListItem() (core.Node, error) {
	li := ast.Li{
		Content: []core.Node{},
	}

	self.iter.Save()
	node, err := self.parseTask()

	if err == nil && node != nil {
		li.Add(node)
		return li, nil
	}

	self.iter.Revert()
	node, err = self.parseBlock()

	if node == nil || err != nil {
		return li, err
	}

	if paragraph, ok := node.(ast.Paragraph); ok {
		li.Add(paragraph.Content...)
	} else {
		li.Add(node)
	}

	return li, nil
}

func (self *Parser) parseTask() (core.Node, error) {
	task := ast.Task{}
	_, err := self.iter.Consume(LeftBracket, "expected '['")

	if err != nil {
		return task, err
	}

	checked, err := self.iter.Consume(Text, "expected ' ' or 'x'")

	if err != nil {
		return task, err
	}

	if checked.String() != " " && checked.String() != "x" {
		return task, errors.New("expected ' ' or 'x'")
	}

	if checked.String() == "x" {
		task.Checked = true
	}

	_, err = self.iter.Consume(RightBracket, "expected ']'")

	if err != nil {
		return task, err
	}

	space, err := self.iter.Consume(Text, "expected ' '")

	if err != nil {
		return task, err
	}

	if space.String() != " " {
		return task, errors.New("expected ' '")
	}

	label := ""

	for !self.iter.Match(NewLine) {
		node, err := self.parseText()

		if err != nil || node == nil {
			return task, err
		}

		text := node.(ast.Text)
		label += text.Content.String()
	}

	task.Label = label
	return task, nil
}

func (self *Parser) parseNewLine() (core.Node, error) {
	if self.iter.Curr.IsWhitespace() {
		return nil, nil
	}

	return ast.NewLine{}, nil
}

func (self *Parser) parseText() (core.Node, error) {
	if self.iter.Curr.Kind == Eof {
		return nil, nil
	}

	node := ast.Text{Content: self.iter.Curr}
	self.iter.Next()
	return node, nil
}
