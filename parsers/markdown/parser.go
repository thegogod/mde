package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/parsers/markdown/ast"
)

type Parser struct {
	curr            *Token
	prev            *Token
	scanner         *Scanner
	blockQuoteDepth int
}

func Markdown() *Parser {
	return &Parser{}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	self.curr = nil
	self.prev = nil
	self.blockQuoteDepth = 0
	self.scanner = NewScanner(src)

	if !self.next() {
		return nil, nil
	}

	group := ast.Group{Items: []core.Node{}}

	for {
		if self.curr.Kind == Eof {
			break
		}

		if self.prev != nil {
			group.Add(ast.NewLine{})
		}

		node, err := self.parseBlock()

		if node == nil {
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
	for range self.blockQuoteDepth {
		if !self.match(BlockQuote) {
			break
		}
	}

	if self.match(Eof) {
		return nil, nil
	} else if self.match(H1) {
		return self.parseHeading(1)
	} else if self.match(H2) {
		return self.parseHeading(2)
	} else if self.match(H3) {
		return self.parseHeading(3)
	} else if self.match(H4) {
		return self.parseHeading(4)
	} else if self.match(H5) {
		return self.parseHeading(5)
	} else if self.match(H6) {
		return self.parseHeading(6)
	} else if self.match(Hr) {
		return self.parseHr()
	} else if self.match(CodeBlock) {
		return self.parseCodeBlock()
	} else if self.match(BlockQuote) {
		return self.parseBlockQuote()
	} else if self.match(NewLine) {
		return self.parseBlock()
	}

	return self.parseParagraph()
}

func (self *Parser) parseInline() (core.Node, error) {
	if self.match(Eof) {
		return nil, nil
	} else if self.match(Bold) {
		return self.parseBold()
	} else if self.match(BoldAlt) {
		return self.parseBoldAlt()
	} else if self.match(Italic) {
		return self.parseItalic()
	} else if self.match(ItalicAlt) {
		return self.parseItalicAlt()
	} else if self.match(Strike) {
		return self.parseStrike()
	} else if self.match(StrikeAlt) {
		return self.parseStrikeAlt()
	} else if self.match(Br) {
		return self.parseBr()
	} else if self.match(Code) {
		return self.parseCode()
	} else if self.match(NewLine) {
		if self.match(NewLine) {
			return nil, nil
		}

		for range self.blockQuoteDepth {
			if !self.match(BlockQuote) {
				break
			}
		}

		return self.parseNewLine()
	}

	return self.parseText()
}

//
// Blocks
//

func (self *Parser) parseHeading(depth int) (core.Node, error) {
	heading := ast.Heading{
		Depth:   depth,
		Content: []core.Node{},
	}

	for self.curr.Kind.IsInline() {
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

	for self.curr.Kind.IsInline() {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return paragraph, err
		}

		paragraph.Add(node)
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

	for !self.match(CodeBlock) {
		node, err := self.parseInline()

		if node == nil || err != nil {
			return code, err
		}

		code.Add(node)
	}

	return code, nil
}

func (self *Parser) parseBlockQuote() (core.Node, error) {
	self.blockQuoteDepth++
	blockQuote := ast.BlockQuote{Content: []core.Node{}}

	for {
		node, err := self.parseBlock()

		if node == nil || err != nil {
			return nil, err
		}

		blockQuote.Add(node)

		if self.curr.Kind != BlockQuote {
			break
		}
	}

	self.blockQuoteDepth--
	return blockQuote, nil
}

//
// InLines
//

func (self *Parser) parseBold() (core.Node, error) {
	bold := ast.Bold{
		Content: []core.Node{},
	}

	for !self.match(Bold) {
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

	for !self.match(BoldAlt) {
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

	for !self.match(Italic) {
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

	for !self.match(ItalicAlt) {
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

	for !self.match(Strike) {
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

	for !self.match(StrikeAlt) {
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

	for !self.match(Code) {
		node, err := self.parseText()

		if node == nil || err != nil {
			return code, err
		}

		code.Add(node)
	}

	return code, nil
}

func (self *Parser) parseNewLine() (core.Node, error) {
	if self.curr.IsWhitespace() {
		return nil, nil
	}

	return ast.NewLine{}, nil
}

func (self *Parser) parseText() (core.Node, error) {
	if self.curr.Kind == Eof {
		return nil, nil
	}

	node := ast.Text{Content: self.curr}
	self.next()
	return node, nil
}

func (self *Parser) next() bool {
	self.prev = self.curr
	token, err := self.scanner.Scan()

	if err != nil {
		return self.next()
	}

	self.curr = token.(*Token)

	if TokenKind(token.GetKind()) == Eof {
		return false
	}

	return true
}

func (self *Parser) match(kind TokenKind) bool {
	if self.curr.Kind != kind {
		return false
	}

	self.next()
	return true
}
