package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/parsers/markdown/ast"
)

type Parser struct {
	curr    *Token
	prev    *Token
	scanner *Scanner
}

func Markdown() *Parser {
	return &Parser{}
}

func (self *Parser) Parse(src []byte) (core.Node, error) {
	self.curr = nil
	self.prev = nil
	self.scanner = NewScanner(src)

	if !self.next() {
		return nil, nil
	}

	group := ast.Group{Items: []core.Node{}}

	for {
		if self.curr.Kind == Eof {
			break
		}

		node, err := self.parseBlock()

		if node == nil {
			continue
		}

		if err != nil {
			return nil, err
		}

		group.Add(node)
		group.Add(ast.NewLine{})
	}

	return group, nil
}

func (self *Parser) parseBlock() (core.Node, error) {
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
	} else if self.match(NewLine) {
		if self.match(NewLine) {
			return nil, nil
		}

		return self.parseNewLine()
	}

	return self.parseText()
}

//
// Blocks
//

func (self *Parser) parseHeading(depth int) (ast.Heading, error) {
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

func (self *Parser) parseParagraph() (ast.Paragraph, error) {
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

func (self *Parser) parseHr() (ast.Hr, error) {
	return ast.Hr{}, nil
}

//
// InLines
//

func (self *Parser) parseBold() (ast.Bold, error) {
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

func (self *Parser) parseBoldAlt() (ast.Bold, error) {
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

func (self *Parser) parseItalic() (ast.Italic, error) {
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

func (self *Parser) parseItalicAlt() (ast.Italic, error) {
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

func (self *Parser) parseStrike() (ast.Strike, error) {
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

func (self *Parser) parseStrikeAlt() (ast.Strike, error) {
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

func (self *Parser) parseBr() (ast.Br, error) {
	return ast.Br{}, nil
}

func (self *Parser) parseNewLine() (ast.NewLine, error) {
	return ast.NewLine{}, nil
}

func (self *Parser) parseText() (ast.Text, error) {
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
