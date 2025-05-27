package markdown

import (
	"unicode"

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

func (self *Scanner) Scan() (core.Token, error) {
	return self.scan()
}

func (self *Scanner) scan() (*Token, error) {
	if self.pos.End >= len(self.src) {
		return self.create(Eof), nil
	}

	self.pos.Start = self.pos.End
	b := self.src[self.pos.End]
	self.next()

	switch b {
	case ' ':
		if self.peek() == ' ' {
			return self.create(Br), nil
		}

		break
	case '\n':
		return self.create(NewLine), nil
	case '#':
		token, err := self.scanHeading()

		if err == nil {
			return token, nil
		}

		break
	case '-':
		if self.peek() == ' ' {
			self.next()
			return self.create(Ul), nil
		} else if self.peek() == '-' {
			self.next()

			if self.peek() == '-' {
				self.next()

				if self.peek() == 0 || unicode.IsSpace(rune(self.peek())) {
					return self.create(Hr), nil
				}
			}
		}

		break
	case '_':
		if self.peek() == '_' {
			self.next()
			return self.create(BoldAlt), nil
		}

		return self.create(ItalicAlt), nil
	case '*':
		if self.peek() == '*' {
			self.next()
			return self.create(Bold), nil
		}

		return self.create(Italic), nil
	case '~':
		if self.peek() == '~' {
			self.next()
			return self.create(StrikeAlt), nil
		}

		return self.create(Strike), nil
	case '>':
		if self.peek() == ' ' {
			self.next()
		}

		return self.create(BlockQuote), nil
	case '`':
		for self.peek() == '`' {
			self.next()
		}

		i := self.pos.End - self.pos.Start

		if i == 1 {
			return self.create(Code), nil
		} else if i >= 3 {
			return self.create(CodeBlock), nil
		}

		break
	case '!':
		return self.create(Bang), nil
	case '[':
		return self.create(LeftBracket), nil
	case ']':
		return self.create(RightBracket), nil
	case '(':
		return self.create(LeftParen), nil
	case ')':
		return self.create(RightParen), nil
	default:
		if unicode.IsNumber(rune(b)) && self.peek() == '.' {
			self.next()

			if self.peek() == ' ' {
				self.next()
				return self.create(Ol), nil
			}

			self.pos.End--
		}
	}

	return self.create(Text), nil
}

func (self *Scanner) scanHeading() (*Token, error) {
	i := 1

	for self.peek() == '#' {
		i++
		self.next()
	}

	if self.peek() != ' ' {
		return nil, self.error("expected space")
	}

	self.next()
	TokenKind := H1

	switch i {
	case 1:
		TokenKind = H1
		break
	case 2:
		TokenKind = H2
		break
	case 3:
		TokenKind = H3
		break
	case 4:
		TokenKind = H4
		break
	case 5:
		TokenKind = H5
		break
	case 6:
		TokenKind = H6
		break
	default:
		return nil, self.error("max heading depth is 6")
	}

	return self.create(TokenKind), nil
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

func (self Scanner) error(message string) error {
	return NewError(
		self.pos,
		message,
	)
}
