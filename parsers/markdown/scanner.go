package markdown

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

func (self *Scanner) Scan() (*core.Token, error) {
	if self._pointer.End.Index >= len(self._pointer.Src) {
		return self.create(core.Eof), nil
	}

	self._pointer.Start = self._pointer.End
	b := self._pointer.Peek()
	self._pointer.Next()

	switch b {
	case ' ':
		return self.create(core.Space), nil
	case '\n':
		return self.create(core.NewLine), nil
	case '\t':
		return self.create(core.Tab), nil
	case '#':
		return self.create(core.Hash), nil
	case '@':
		return self.create(core.At), nil
	case '.':
		return self.create(core.Period), nil
	case '|':
		return self.create(core.Pipe), nil
	case '&':
		return self.create(core.Ampersand), nil
	case '-':
		return self.create(core.Dash), nil
	case '_':
		return self.create(core.Underscore), nil
	case '*':
		return self.create(core.Asterisk), nil
	case '~':
		return self.create(core.Tilde), nil
	case '=':
		return self.create(core.Equals), nil
	case '>':
		if self._pointer.Peek() == ' ' {
			self._pointer.Next()
		}

		return self.create(core.GreaterThan), nil
	case '<':
		if self._pointer.Peek() == ' ' {
			self._pointer.Next()
		}

		return self.create(core.LessThan), nil
	case '`':
		return self.create(core.BackQuote), nil
	case ':':
		return self.create(core.Colon), nil
	case '!':
		return self.create(core.Bang), nil
	case '[':
		return self.create(core.LeftBracket), nil
	case ']':
		return self.create(core.RightBracket), nil
	case '{':
		return self.create(core.LeftBrace), nil
	case '}':
		return self.create(core.RightBrace), nil
	case '(':
		return self.create(core.LeftParen), nil
	case ')':
		return self.create(core.RightParen), nil
	case '/':
		return self.create(core.Slash), nil
	case '\\':
		return self.create(core.BackSlash), nil
	default:
		if b >= '0' && b <= '9' {
			return self.scanNumeric()
		}
	}

	return self.scanText()
}

func (self *Scanner) scanNumeric() (*core.Token, error) {
	for self._pointer.Peek() >= '0' && self._pointer.Peek() <= '9' {
		self._pointer.Next()
	}

	if self._pointer.Peek() != '.' {
		return self.create(core.Integer), nil
	}

	self.Save()
	self._pointer.Next()

	if self._pointer.Peek() < '0' || self._pointer.Peek() > '9' {
		self.RevertAndPop()
		return self.create(core.Integer), nil
	}

	for self._pointer.Peek() >= '0' && self._pointer.Peek() <= '9' {
		self._pointer.Next()
	}

	self.Pop()
	return self.create(core.Decimal), nil
}

func (self *Scanner) scanText() (*core.Token, error) {
	for (self._pointer.Peek() >= 'a' && self._pointer.Peek() <= 'z') || (self._pointer.Peek() >= 'A' && self._pointer.Peek() <= 'Z') {
		self._pointer.Next()
	}

	return self.create(core.Text), nil
}

func (self Scanner) create(kind core.TokenKind) *core.Token {
	return core.NewToken(
		kind,
		self._pointer.Start,
		self._pointer.End,
		self._pointer.Bytes(),
	)
}
