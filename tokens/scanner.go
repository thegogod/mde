package tokens

import (
	"github.com/thegogod/mde/core"
)

type Scanner struct {
	src    []byte
	start  core.Position
	end    core.Position
	_saves []core.Position
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		src:    src,
		start:  core.Position{},
		end:    core.Position{},
		_saves: []core.Position{},
	}
}

func (self *Scanner) Save() {
	if self._saves == nil {
		self._saves = []core.Position{}
	}

	self._saves = append(self._saves, self.end)
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

	self.end = self._saves[len(self._saves)-1]
}

func (self *Scanner) RevertAndPop() {
	self.Revert()
	self.Pop()
}

func (self *Scanner) Scan() (core.Token, error) {
	if self.end.Index >= len(self.src) {
		return self.create(Eof), nil
	}

	self.start = self.end
	b := self.src[self.end.Index]
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
		if self.peek() == ' ' {
			self.next()
		}

		return self.create(GreaterThan), nil
	case '<':
		if self.peek() == ' ' {
			self.next()
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

func (self *Scanner) scanText() (*Token, error) {
	for (self.peek() >= 'a' && self.peek() <= 'z') || (self.peek() >= 'A' && self.peek() <= 'Z') {
		self.next()
	}

	return self.create(Text), nil
}

func (self *Scanner) next() byte {
	self.end.Index++
	self.end.Col++

	if self.peek() == '\n' {
		self.end.Ln++
		self.end.Col = 0
	}

	return self.peek()
}

func (self Scanner) peek() byte {
	if self.end.Index >= len(self.src) {
		return byte(Eof)
	}

	return self.src[self.end.Index]
}

func (self Scanner) create(TokenKind TokenKind) *Token {
	return NewToken(
		TokenKind,
		self.end,
		self.src[self.start.Index:self.end.Index],
	)
}
