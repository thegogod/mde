package markdown

import (
	"unicode"

	mde "github.com/thegogod/mde/core"
)

type Scanner struct {
	src   []byte
	ln    int
	left  int
	right int
}

func NewScanner(src []byte) *Scanner {
	return &Scanner{
		src:   src,
		ln:    0,
		left:  0,
		right: 0,
	}
}

func (self *Scanner) Scan() (mde.Token, error) {
	if self.right >= len(self.src) {
		return self.create(Eof), nil
	}

	self.left = self.right
	b := self.src[self.right]
	self.right++

	switch b {
	case ' ':
		if self.peek() == ' ' {
			return self.create(Br), nil
		}

		break
	case '\n':
		return self.create(Space), nil
	case '#':
		token, err := self.onHeading()

		if err == nil {
			return token, nil
		}

		break
	case '-':
		if self.peek() == ' ' {
			self.right++
			return self.create(Ul), nil
		} else if self.peek() == '-' {
			self.right++

			if self.peek() == '-' {
				self.right++

				if self.peek() == 0 || unicode.IsSpace(rune(self.peek())) {
					return self.create(Hr), nil
				}
			}
		}

		break
	case '_':
		if self.peek() == '_' {
			self.right++
			return self.create(BoldAlt), nil
		}

		return self.create(ItalicAlt), nil
	case '*':
		if self.peek() == '*' {
			self.right++
			return self.create(Bold), nil
		}

		return self.create(Italic), nil
	case '>':
		return self.create(BlockQuote), nil
	case '`':
		for self.peek() == '`' {
			self.right++
		}

		i := self.right - self.left

		if i == 1 {
			return self.create(Code), nil
		} else if i == 3 {
			return self.create(CodeBlock), nil
		}

		break
	default:
		if unicode.IsNumber(rune(b)) && self.peek() == '.' {
			self.right++

			if self.peek() == ' ' {
				return self.create(Ol), nil
			}

			self.right--
		}
	}

	return self.create(Text), nil
}

func (self *Scanner) onHeading() (*Token, error) {
	i := 1

	for self.peek() == '#' {
		i++
		self.right++
	}

	if self.peek() != ' ' {
		return nil, self.error("expected space")
	}

	self.right++
	self.left = self.right

	for self.right < len(self.src) && self.peek() != '\n' {
		self.right++
	}

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

func (self Scanner) peek() byte {
	if self.right >= len(self.src) {
		return byte(Eof)
	}

	return self.src[self.right]
}

func (self Scanner) create(TokenKind TokenKind) *Token {
	return NewToken(
		TokenKind,
		mde.Position{
			Ln:    self.ln,
			Start: self.left,
			End:   self.right,
		},
		self.src[self.left:self.right],
	)
}

func (self Scanner) error(message string) error {
	return NewError(
		mde.Position{
			Ln:    self.ln,
			Start: self.left,
			End:   self.right,
		},
		message,
	)
}
