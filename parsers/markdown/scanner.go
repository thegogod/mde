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
	if self._pointer.Eof() {
		return self._pointer.Done(core.Eof), nil
	}

	self._pointer.Start = self._pointer.End
	b := self._pointer.Peek()
	self._pointer.Next()

	switch b {
	case ' ':
		return self._pointer.Done(core.Space), nil
	case '\n':
		return self._pointer.Done(core.NewLine), nil
	case '\t':
		return self._pointer.Done(core.Tab), nil
	case '#':
		return self._pointer.Done(core.Hash), nil
	case '@':
		return self._pointer.Done(core.At), nil
	case '.':
		return self._pointer.Done(core.Period), nil
	case '|':
		return self._pointer.Done(core.Pipe), nil
	case '&':
		return self._pointer.Done(core.Ampersand), nil
	case '-':
		return self._pointer.Done(core.Dash), nil
	case '_':
		return self._pointer.Done(core.Underscore), nil
	case '*':
		return self._pointer.Done(core.Asterisk), nil
	case '~':
		return self._pointer.Done(core.Tilde), nil
	case '=':
		return self._pointer.Done(core.Equals), nil
	case '>':
		if self._pointer.Peek() == ' ' {
			self._pointer.Next()
		}

		return self._pointer.Done(core.GreaterThan), nil
	case '<':
		if self._pointer.Peek() == ' ' {
			self._pointer.Next()
		}

		return self._pointer.Done(core.LessThan), nil
	case '`':
		return self._pointer.Done(core.BackQuote), nil
	case ':':
		return self._pointer.Done(core.Colon), nil
	case '!':
		return self._pointer.Done(core.Bang), nil
	case '[':
		return self._pointer.Done(core.LeftBracket), nil
	case ']':
		return self._pointer.Done(core.RightBracket), nil
	case '{':
		return self._pointer.Done(core.LeftBrace), nil
	case '}':
		return self._pointer.Done(core.RightBrace), nil
	case '(':
		return self._pointer.Done(core.LeftParen), nil
	case ')':
		return self._pointer.Done(core.RightParen), nil
	case '/':
		return self._pointer.Done(core.Slash), nil
	case '\\':
		return self._pointer.Done(core.BackSlash), nil
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
		return self._pointer.Done(core.Integer), nil
	}

	self.Save()
	self._pointer.Next()

	if self._pointer.Peek() < '0' || self._pointer.Peek() > '9' {
		self.RevertAndPop()
		return self._pointer.Done(core.Integer), nil
	}

	for self._pointer.Peek() >= '0' && self._pointer.Peek() <= '9' {
		self._pointer.Next()
	}

	self.Pop()
	return self._pointer.Done(core.Decimal), nil
}

func (self *Scanner) scanText() (*core.Token, error) {
	for (self._pointer.Peek() >= 'a' && self._pointer.Peek() <= 'z') || (self._pointer.Peek() >= 'A' && self._pointer.Peek() <= 'Z') {
		self._pointer.Next()
	}

	return self._pointer.Done(core.Text), nil
}
