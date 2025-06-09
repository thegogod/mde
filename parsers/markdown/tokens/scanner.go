package tokens

import (
	"github.com/thegogod/mde/core"
)

type Scanner struct {
	src []byte
	pos core.Position
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		src: src,
		pos: core.Position{},
	}
}

func (self *Scanner) Save() {
	self.pos.Save()
}

func (self *Scanner) Pop() {
	self.pos.Pop()
}

func (self *Scanner) Revert() {
	self.pos.Revert()
}

func (self *Scanner) RevertAndPop() {
	self.pos.RevertAndPop()
}

func (self *Scanner) Scan() (core.Token, error) {
	if self.pos.End >= len(self.src) {
		return self.create(Eof), nil
	}

	self.pos.Start = self.pos.End
	b := self.src[self.pos.End]
	self.next()

	switch b {
	case ' ':
		return self.create(Space), nil
	case '\n':
		return self.create(NewLine), nil
	case '\t':
		return self.create(Tab), nil
	case '#':
		return self.create(Hash), nil
	case '.':
		return self.create(Period), nil
	case '|':
		return self.create(Pipe), nil
	case '&':
		return self.create(Ampersand), nil
	case '-':
		return self.create(Dash), nil
	case '_':
		return self.create(Underscore), nil
	case '*':
		return self.create(Asterisk), nil
	case '~':
		return self.create(Tilde), nil
	case '=':
		return self.create(Equals), nil
	case '>':
		if self.peek() == ' ' {
			self.next()
		}

		return self.create(GreaterThan), nil
	case '`':
		return self.create(BackQuote), nil
	case ':':
		return self.create(Colon), nil
	case '!':
		return self.create(Bang), nil
	case '[':
		return self.create(LeftBracket), nil
	case ']':
		return self.create(RightBracket), nil
	case '{':
		return self.create(LeftBrace), nil
	case '}':
		return self.create(RightBrace), nil
	case '(':
		return self.create(LeftParen), nil
	case ')':
		return self.create(RightParen), nil
	default:
		if b >= '0' && b <= '9' {
			return self.scanNumeric()
		}
	}

	return self.create(Text), nil
}

func (self *Scanner) scanNumeric() (*Token, error) {
	for self.peek() >= '0' && self.peek() <= '9' {
		self.next()
	}

	if self.peek() != '.' {
		return self.create(Integer), nil
	}

	self.Save()
	self.next()

	if self.peek() < '0' || self.peek() > '9' {
		self.RevertAndPop()
		return self.create(Integer), nil
	}

	for self.peek() >= '0' && self.peek() <= '9' {
		self.next()
	}

	self.Pop()
	return self.create(Decimal), nil
}

func (self *Scanner) next() byte {
	self.pos.End++
	self.pos.Col++

	if self.peek() == '\n' {
		self.pos.Ln++
		self.pos.Col = 0
	}

	return self.peek()
}

func (self Scanner) peek() byte {
	if self.pos.End >= len(self.src) {
		return byte(Eof)
	}

	return self.src[self.pos.End]
}

func (self Scanner) create(TokenKind TokenKind) *Token {
	return NewToken(
		TokenKind,
		self.pos,
		self.src[self.pos.Start:self.pos.End],
	)
}
