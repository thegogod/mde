package tokens

import (
	"github.com/thegogod/mde/core"
)

type Scanner struct {
	_pointer *core.Pointer
	_saves   []core.Position
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		_pointer: core.Ptr(src),
		_saves:   []core.Position{},
	}
}

func (self Scanner) Pointer() *core.Pointer {
	return self._pointer
}

func (self *Scanner) Save() {
	if self._saves == nil {
		self._saves = []core.Position{}
	}

	self._saves = append(self._saves, self._pointer.End)
}

func (self *Scanner) Pop() {
	if len(self._saves) == 0 {
		return
	}

	self._saves = self._saves[:len(self._saves)-1]
}

func (self *Scanner) Revert() {
	if self._saves == nil || len(self._saves) == 0 {
		return
	}

	self._pointer.End = self._saves[len(self._saves)-1]
}

func (self *Scanner) RevertAndPop() {
	self.Revert()
	self.Pop()
}

func (self *Scanner) Scan() (core.Token, error) {
	if self._pointer.End.Index >= len(self._pointer.Src) {
		return self.create(Eof), nil
	}

	self._pointer.Start = self._pointer.End
	b := self._pointer.Peek()
	self._pointer.Next()

	switch b {
	case ' ':
		return self.create(Space), nil
	case '\n':
		return self.create(NewLine), nil
	case '\t':
		return self.create(Tab), nil
	case '#':
		return self.create(Hash), nil
	case '@':
		return self.create(At), nil
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
		if self._pointer.Peek() == ' ' {
			self._pointer.Next()
		}

		return self.create(GreaterThan), nil
	case '<':
		if self._pointer.Peek() == ' ' {
			self._pointer.Next()
		}

		return self.create(LessThan), nil
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
	case '/':
		return self.create(Slash), nil
	case '\\':
		return self.create(BackSlash), nil
	default:
		if b >= '0' && b <= '9' {
			return self.scanNumeric()
		}
	}

	return self.scanText()
}

func (self *Scanner) scanNumeric() (*Token, error) {
	for self._pointer.Peek() >= '0' && self._pointer.Peek() <= '9' {
		self._pointer.Next()
	}

	if self._pointer.Peek() != '.' {
		return self.create(Integer), nil
	}

	self.Save()
	self._pointer.Next()

	if self._pointer.Peek() < '0' || self._pointer.Peek() > '9' {
		self.RevertAndPop()
		return self.create(Integer), nil
	}

	for self._pointer.Peek() >= '0' && self._pointer.Peek() <= '9' {
		self._pointer.Next()
	}

	self.Pop()
	return self.create(Decimal), nil
}

func (self *Scanner) scanText() (*Token, error) {
	for (self._pointer.Peek() >= 'a' && self._pointer.Peek() <= 'z') || (self._pointer.Peek() >= 'A' && self._pointer.Peek() <= 'Z') {
		self._pointer.Next()
	}

	return self.create(Text), nil
}

func (self Scanner) create(TokenKind TokenKind) *Token {
	return NewToken(
		TokenKind,
		self._pointer.Start,
		self._pointer.End,
		self._pointer.Bytes(),
	)
}
